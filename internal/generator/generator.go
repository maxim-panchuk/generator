package generator

import (
	"fmt"
	"generator/internal/generator/definitions"
	"generator/internal/generator/folderstructure"
	openapiParser "generator/internal/generator/parser"
	"generator/internal/generator/templates"
	"generator/internal/generator/templates/models"
	"generator/internal/logger"
	"path"
)

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) Generate() error {
	parser, err := openapiParser.New()
	if err != nil {
		return fmt.Errorf("generate: %e", err)
	}

	if err := parser.Parse(); err != nil {
		return fmt.Errorf("generate: %e", err)
	}

	if err := g.generateModels(); err != nil {
		return fmt.Errorf("generate: %e", err)
	}

	return nil
}

func (g *Generator) generateModels() error {
	for modelName, model := range definitions.GetData().Models {
		logger.Info(fmt.Sprintf("generate model: %s", modelName))
		if err := templates.RunTemplate(&templates.TemplateData{
			Template: models.Models(),
			FilePath: path.Join(folderstructure.PathToModel, model.ModelName+".go"),
			Data:     model,
		}); err != nil {
			return fmt.Errorf("generate models: %e", err)
		}
	}
	return nil
}
