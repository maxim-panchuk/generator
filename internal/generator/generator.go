package generator

import (
	"fmt"
	"generator/internal/generator/crudgen"
	"generator/internal/generator/definitions"
	"generator/internal/generator/filesystem"
	openapiParser "generator/internal/generator/parser"
	"generator/internal/generator/templates"
	"generator/internal/generator/templates/common"
	"generator/internal/generator/templates/mapper"
	"generator/internal/generator/templates/models"
	"generator/internal/generator/utils"
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

	data := definitions.GetData()
	logger.Info(fmt.Sprintf("data %v", data))

	if err := filesystem.CreateArchetypeFileSystem(); err != nil {
		return fmt.Errorf("generate: %e", err)
	}

	if err := g.generateMain(); err != nil {
		return fmt.Errorf("generate: %e", err)
	}

	if err := g.generateModels(); err != nil {
		return fmt.Errorf("generate: %e", err)
	}

	if err := g.generateCruds(); err != nil {
		return fmt.Errorf("generate: %e", err)
	}

	if err := g.generateMappers(); err != nil {
		return fmt.Errorf("generate: %e", err)
	}

	return nil
}

func (g *Generator) generateMain() error {
	if err := templates.RunTemplate(&templates.TemplateData{
		Template: common.Main(),
		FilePath: path.Join(filesystem.PathToApp, "main.go"),
		Data:     definitions.GetData(),
	}); err != nil {
		return fmt.Errorf("generateMain: %e", err)
	}
	return nil
}

func (g *Generator) generateModels() error {
	for modelName, model := range definitions.GetData().Models {
		logger.Info(fmt.Sprintf("generate model: %s", modelName))
		if err := filesystem.CreateDir(path.Join(filesystem.PathToModel, utils.LowFirst(model.ModelName))); err != nil {
			return fmt.Errorf("generate models: %e", err)
		}
		if err := templates.RunTemplate(&templates.TemplateData{
			Template: models.Models(),
			FilePath: path.Join(filesystem.PathToModel, utils.LowFirst(model.ModelName), model.ModelName+".go"),
			Data:     model,
		}); err != nil {
			return fmt.Errorf("generate models: %e", err)
		}
	}
	return nil
}

func (g *Generator) generateCruds() error {
	crudGen := crudgen.New()
	if err := crudGen.Generate(); err != nil {
		return fmt.Errorf("generateCruds: %e", err)
	}
	return nil
}

func (g *Generator) generateMappers() error {
	for modelName, model := range definitions.GetData().Models {
		if model.XDb == nil {
			continue
		}
		logger.Info(fmt.Sprintf("generate mapper: %s", modelName))
		if err := filesystem.CreateDir(path.Join(filesystem.PathToMapper, utils.LowFirst(model.ModelName))); err != nil {
			return fmt.Errorf("generate mappers: %e", err)
		}
		if err := templates.RunTemplate(&templates.TemplateData{
			Template: mapper.Mapper(),
			FilePath: path.Join(filesystem.PathToMapper, utils.LowFirst(model.ModelName), model.ModelName+".go"),
			Data:     model,
		}); err != nil {
			return fmt.Errorf("generate mappers: %e", err)
		}
	}
	return nil
}
