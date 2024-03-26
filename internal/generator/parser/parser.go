package parser

import (
	"fmt"
	"generator/config"
	"generator/internal/generator/definitions"
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
	return nil
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
