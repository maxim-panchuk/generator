package parser

import (
	"fmt"
	"generator/config"
	"generator/internal/generator/definitions"
	"generator/internal/generator/utils"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"os"
)

type Parser struct {
	doc *libopenapi.DocumentModel[v3.Document]
}

func New() (*Parser, error) {
	doc, err := parseOpenApiV3(config.Get().PathToSwaggerFile)
	if err != nil {
		return nil, fmt.Errorf("new: %e", err)
	}
	return &Parser{doc: doc}, nil
}

func (p *Parser) Parse() error {
	if err := p.addModels(); err != nil {
		return fmt.Errorf("parse: %e", err)
	}
	if err := p.addTags(); err != nil {
		return fmt.Errorf("parse: %e", err)
	}
	return nil
}

func (p *Parser) addTags() error {
	tags := definitions.GetData().Tags
	for _, tag := range p.doc.Model.Tags {
		pathList, err := p.iterateUris(tag.Name)
		if err != nil {
			return fmt.Errorf("addTags: %s", tag.Name)
		}
		tags[tag.Name] = pathList
	}
	definitions.GetData().Tags = tags
	return nil
}

func (p *Parser) iterateUris(tag string) ([]*definitions.Path, error) {
	pathList := make([]*definitions.Path, 0)
	for pair := p.doc.Model.Paths.PathItems.First(); pair != nil; pair = pair.Next() {
		url := pair.Key()
		operations, err := p.iterateCruds(tag, url, pair.Value().GetOperations())
		if err != nil {
			return nil, fmt.Errorf("iterateUris: %s", url)
		}
		if len(operations) == 0 {
			continue
		}
		path := &definitions.Path{
			Url:        url,
			Operations: operations,
		}
		pathList = append(pathList, path)
	}
	return pathList, nil
}

func (p *Parser) iterateCruds(tag, url string, cruds *orderedmap.Map[string, *v3.Operation]) ([]*definitions.Operation, error) {
	operationList := make([]*definitions.Operation, 0)
	for pair := cruds.First(); pair != nil; pair = pair.Next() {
		opv3 := pair.Value()
		if tag != opv3.Tags[0] {
			continue
		}
		op, err := p.buildOperation(opv3)
		op.Url = url
		op.Type = pair.Key()
		if err != nil {
			return nil, fmt.Errorf("iterateCruds: %s", pair.Key())
		}
		operationList = append(operationList, op)
	}
	return operationList, nil
}

func (p *Parser) buildOperation(opv3 *v3.Operation) (*definitions.Operation, error) {
	paramList, err := parseParamList(opv3)
	if err != nil {
		return nil, fmt.Errorf("buildOperation: %s", opv3.OperationId)
	}

	op := &definitions.Operation{
		Tag:         opv3.Tags[0],
		Summary:     opv3.Summary,
		Description: opv3.Description,
		OperationId: opv3.OperationId,
		Parameters:  paramList,
		RequestBody: parseRequestBody(opv3),
		Responses:   parseResponses(opv3),
		IsTypical:   false,
	}

	xm := extractXMeta(opv3)
	if xm != nil {
		xm.IncludeModel(utils.GetDefinitionNameFromRef(xm.Object))
		op.XMeta = xm
		op.IsTypical = true
	}
	return op, nil
}

func parseParamList(opv3 *v3.Operation) ([]*definitions.Parameter, error) {
	pv3List := opv3.Parameters
	if pv3List == nil {
		return nil, nil
	}
	paramList := make([]*definitions.Parameter, 0, len(pv3List))
	for _, pv3 := range pv3List {
		p, err := parseParam(pv3)
		if err != nil {
			return nil, fmt.Errorf("parseParamList ")
		}
		paramList = append(paramList, p)
	}
	return paramList, nil
}

func parseParam(pv3 *v3.Parameter) (*definitions.Parameter, error) {
	p := &definitions.Parameter{
		Name:        pv3.Name,
		Description: pv3.Description,
		In:          pv3.In,
		IsArray:     false,
		Required:    true,
	}

	if required := pv3.Required; required != nil {
		p.Required = *required
	}

	schema, err := buildSchema(pv3.Schema)
	if err != nil {
		return nil, fmt.Errorf("parse param: %s", pv3.Name)
	}
	isArray, arraySchema := utils.IsArraySchema(schema)
	if isArray {
		p.IsArray = true
		schema = arraySchema.Schema()
	}
	p.Type = schema.Type[0]
	p.Format = schema.Format
	if def := schema.Default; def != nil {
		p.Default = schema.Default.Value
	}
	return p, nil
}

func parseRequestBody(opv3 *v3.Operation) *definitions.RequestBody {
	if opv3.RequestBody == nil || opv3.RequestBody.Content == nil {
		return nil
	}
	content, isArray := parseContent(opv3.RequestBody.Content)
	if content == nil {
		return nil
	}
	return &definitions.RequestBody{
		Description: opv3.RequestBody.Description,
		Content:     content,
		IsArray:     isArray,
	}
}

func parseResponses(opv3 *v3.Operation) []*definitions.Response {
	responseList := make([]*definitions.Response, 0)
	for pair := opv3.Responses.Codes.First(); pair != nil; pair = pair.Next() {
		code := pair.Key()
		resp := pair.Value()
		content, isArray := parseContent(resp.Content)
		readyResp := &definitions.Response{
			Code:        code,
			Description: resp.Description,
			Content:     content,
			IsArray:     isArray,
		}

		responseList = append(responseList, readyResp)
	}
	if def := parseDefault(opv3); def != nil {
		responseList = append(responseList, def)
	}
	return responseList
}

func parseDefault(opv3 *v3.Operation) *definitions.Response {
	def := opv3.Responses.Default
	if def == nil {
		return nil
	}
	content, isArray := parseContent(def.Content)
	readyResp := &definitions.Response{
		Code:        "default",
		Description: def.Description,
		Content:     content,
		IsArray:     isArray,
	}
	return readyResp
}

func parseContent(content *orderedmap.Map[string, *v3.MediaType]) (*definitions.Model, bool) {
	if content == nil {
		return nil, false
	}
	c, ok := content.Get("application/json")
	if !ok {
		return nil, false
	}
	schema, err := c.Schema.BuildSchema()
	if err != nil {
		panic(err)
	}
	var modelName string
	isArray, arraySchema := utils.IsArraySchema(schema)
	if isArray {
		modelName = utils.GetDefinitionNameFromRef(arraySchema.GetReference())
	} else {
		modelName = utils.GetDefinitionNameFromRef(c.Schema.GetReference())
	}
	model := definitions.GetData().Models[modelName]
	return model, isArray
}

func extractXMeta(op *v3.Operation) *definitions.XMeta {
	node, ok := op.Extensions.Get("x-meta")
	if !ok {
		return nil
	}

	xm := &definitions.XMeta{}
	if err := node.Decode(xm); err != nil {
		panic(err)
	}

	return xm
}

func (p *Parser) addModels() error {
	for pair := p.doc.Model.Components.Schemas.First(); pair != nil; pair = pair.Next() {
		model, err := p.parseModel(pair.Value(), pair.Key())
		if err != nil {
			return fmt.Errorf("parse models: %e", err)
		}
		definitions.GetData().Models[model.ModelName] = model
	}
	return nil
}

func (p *Parser) parseModel(schemaProxy *base.SchemaProxy, schemaName string) (*definitions.Model, error) {
	schema, err := buildSchema(schemaProxy)
	if err != nil {
		return nil, fmt.Errorf("parse models: %w", err)
	}

	model := &definitions.Model{
		ModelName:    schemaName,
		Type:         schema.Type[0],
		Format:       schema.Format,
		Description:  schema.Description,
		IsEnum:       false,
		EnumValues:   nil,
		XDb:          extractXDb(schema),
		PostgresType: utils.ConvertToPostgresType(schema.Type[0], schema.Format),
	}

	if schema.Enum != nil && len(schema.Enum) > 0 {
		enumValues := make([]interface{}, 0, len(schema.Enum))
		for _, enum := range schema.Enum {
			enumValues = append(enumValues, enum.Value)
		}
		model.EnumValues = enumValues
		model.IsEnum = true
	}

	if ref := schemaProxy.GetReference(); ref != "" {
		model.Ref = ref
	}
	isArray, arraySchema := utils.IsArraySchema(schema)
	if isArray {
		arraySchemaName := utils.GetDefinitionNameFromRef(arraySchema.GetReference())
		propItemsModel, err := p.parseModel(arraySchema, arraySchemaName)
		if err != nil {
			return nil, fmt.Errorf("parse model %s: %w", arraySchemaName, err)
		}
		model.Items = propItemsModel
	}

	if !containsProperties(schema) {
		return model, nil
	}

	model.PropKeys = make([]string, 0)
	model.Properties = orderedmap.New[string, *definitions.Model]()

	for pair := schema.Properties.First(); pair != nil; pair = pair.Next() {
		propName := pair.Key()
		propSchemaProxy := pair.Value()
		propModel, err := p.parseModel(propSchemaProxy, propName)
		if err != nil {
			return nil, fmt.Errorf("parse model: %e", err)
		}
		model.Properties.Set(propName, propModel)
		model.PropKeys = append(model.PropKeys, propName)
	}
	return model, nil
}

func extractXDb(s *base.Schema) *definitions.XDb {
	node, ok := s.Extensions.Get("x-db")
	if !ok {
		return nil
	}

	xb := &definitions.XDb{}
	if err := node.Decode(xb); err != nil {
		panic(err)
	}

	return xb
}

func containsProperties(schema *base.Schema) bool {
	if schema.Properties == nil {
		return false
	}
	return schema.Properties.Len() > 0
}

func buildSchema(proxy *base.SchemaProxy) (*base.Schema, error) {
	schema, err := proxy.BuildSchema()
	if err != nil {
		return nil, fmt.Errorf("build schema: %w", err)
	}
	return schema, nil
}

func parseOpenApiV3(path string) (*libopenapi.DocumentModel[v3.Document], error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("parse cannot parse document: %e", err)
	}

	document, err := libopenapi.NewDocument(file)

	if err != nil {
		return nil, fmt.Errorf("parse cannot create new document: %e", err)
	}

	v3Model, errors := document.BuildV3Model()

	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from "+
			"document: %d errors reported",
			len(errors)))
	}

	return v3Model, nil
}
