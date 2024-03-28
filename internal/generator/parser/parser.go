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
	p.addTags()
	return nil
}

func (p *Parser) addTags() {
	tags := definitions.GetData().Tags
	for _, tag := range p.doc.Model.Tags {
		tags[tag.Name] = p.iterateUris(tag.Name)
	}
	definitions.GetData().Tags = tags
}

func (p *Parser) iterateUris(tag string) []*definitions.Path {
	pathList := make([]*definitions.Path, 0)
	for pair := p.doc.Model.Paths.PathItems.First(); pair != nil; pair = pair.Next() {
		uri := pair.Key()
		operations := p.iterateCruds(tag, pair.Value().GetOperations())
		if len(operations) == 0 {
			continue
		}
		path := &definitions.Path{
			Url:        uri,
			Operations: operations,
		}
		pathList = append(pathList, path)
	}
	return pathList
}

func (p *Parser) iterateCruds(tag string, cruds *orderedmap.Map[string, *v3.Operation]) []*definitions.Operation {
	operationList := make([]*definitions.Operation, 0)
	for pair := cruds.First(); pair != nil; pair = pair.Next() {
		opv3 := pair.Value()
		if tag != opv3.Tags[0] {
			continue
		}
		operationList = append(operationList, p.buildOperation(opv3))
	}
	return operationList
}

func (p *Parser) buildOperation(opv3 *v3.Operation) *definitions.Operation {
	op := &definitions.Operation{
		Tag:         opv3.Tags[0],
		Summary:     opv3.Summary,
		Description: opv3.Description,
		OperationId: opv3.OperationId,
		// TODO:
		Parameters: nil,
		// TODO:
		RequestBody: nil,
		Responses:   parseResponses(opv3),
		IsTypical:   false,
	}

	xm := extractXMeta(opv3)
	if xm != nil {
		xm.IncludeModel(utils.GetDefinitionNameFromRef(xm.Object))
		op.XMeta = xm
		op.IsTypical = true
	}
	return op
}

func parseResponses(opv3 *v3.Operation) []*definitions.Response {
	responseList := make([]*definitions.Response, 0)
	for pair := opv3.Responses.Codes.First(); pair != nil; pair = pair.Next() {
		code := pair.Key()
		resp := pair.Value()

		readyResp := &definitions.Response{
			Code:        code,
			Description: resp.Description,
			Content:     parseContent(resp.Content),
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
	readyResp := &definitions.Response{
		Code:        "default",
		Description: def.Description,
		Content:     parseContent(def.Content),
	}
	return readyResp
}

func parseContent(content *orderedmap.Map[string, *v3.MediaType]) *definitions.Model {
	if content == nil {
		return nil
	}
	c, ok := content.Get("application/json")
	if !ok {
		return nil
	}
	modelName := utils.GetDefinitionNameFromRef(c.Schema.GetReference())
	return definitions.GetData().Models[modelName]
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
		ModelName:   schemaName,
		Type:        schema.Type[0],
		Format:      schema.Format,
		Description: schema.Description,
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
