package crudgen

import (
	"fmt"
	"generator/config"
	"generator/internal/generator/definitions"
	"generator/internal/generator/filesystem"
	"generator/internal/generator/templates"
	"generator/internal/generator/templates/layer/controller"
	"generator/internal/generator/templates/layer/repository"
	"generator/internal/generator/templates/layer/service"
	"generator/internal/logger"
	"path"
	"text/template"
)

var targetControllerDirPath = path.Join(config.Get().PathToRepositoryRoot, "internal", "transport", "http", "in")
var targetRepositoryDirPath = path.Join(config.Get().PathToRepositoryRoot, "internal", "database", "repositories")
var targetServiceDirPath = path.Join(config.Get().PathToRepositoryRoot, "internal", "service")

type CrudGen struct {
	Data *definitions.Data
}

func New() *CrudGen { return &CrudGen{Data: definitions.GetData()} }

func (g *CrudGen) Generate() error {
	for tag, _ := range g.Data.Tags {
		controllerGenerator := newLayerGenerator(tag, &LayerData{
			Tag:  tag,
			Data: g.Data,
		}).withController()
		if err := controllerGenerator.genLayer(); err != nil {
			return fmt.Errorf("generate: %e", err)
		}

		serviceGenerator := newLayerGenerator(tag, &LayerData{
			Tag:  tag,
			Data: g.Data,
		}).withService()
		logger.Info(fmt.Sprintf("gen service layer for: %s", tag))
		if err := serviceGenerator.genLayer(); err != nil {
			return fmt.Errorf("generate: %e", err)
		}
	}
	return nil
}

type Template interface {
	Interface() *template.Template
	GeneratedImpl() *template.Template
	GeneratedInit() *template.Template
	CustomInit() *template.Template
}

type LayerData struct {
	Tag  string
	Data *definitions.Data
}

type LayerGenerator struct {
	tag                string
	targetLayerDirPath string
	tmpl               Template
	data               *LayerData
}

func newLayerGenerator(tag string, data *LayerData) *LayerGenerator {
	return &LayerGenerator{tag: tag, data: data}
}

func (g *LayerGenerator) withController() *LayerGenerator {
	g.targetLayerDirPath = targetControllerDirPath
	g.tmpl = controller.New()
	return g
}

func (g *LayerGenerator) withService() *LayerGenerator {
	g.targetLayerDirPath = targetServiceDirPath
	g.tmpl = service.New()
	return g
}

func (g *LayerGenerator) withRepository() *LayerGenerator {
	g.targetLayerDirPath = targetRepositoryDirPath
	g.tmpl = repository.New()
	return g
}

func (g *LayerGenerator) genLayer() error {
	if err := g.mkDirStructure(); err != nil {
		return fmt.Errorf("genLayer: %w", err)
	}

	if err := g.genLayerInterface(); err != nil {
		return fmt.Errorf("genLayer: %w", err)
	}

	//if err := g.genLayerDefaultImpl(); err != nil {
	//	return fmt.Errorf("genLayer: %w", err)
	//}

	return nil
}

func (g *LayerGenerator) mkDirStructure() error {
	if err := g.mkBoDir(); err != nil {
		return fmt.Errorf("mkDirStructure: %w", err)
	}

	if err := g.mkdirGenerated(); err != nil {
		return fmt.Errorf("mkDirStructure: %w", err)
	}

	if err := g.mkdirCustom(); err != nil {
		return fmt.Errorf("mkDirStructure: %w", err)
	}

	return nil
}

func (g *LayerGenerator) mkBoDir() error {
	if err := filesystem.CreateDir(g.getPathToDirBo()); err != nil {
		return fmt.Errorf("mkBoDir: %w", err)
	}
	return nil
}

func (g *LayerGenerator) mkdirGenerated() error {
	if err := filesystem.CreateDir(g.getPathToGenerated()); err != nil {
		return fmt.Errorf("mkdirGenerated: %w", err)
	}
	return nil
}

func (g *LayerGenerator) mkdirCustom() error {
	if err := filesystem.CreateDir(g.getPathToCustom()); err != nil {
		return fmt.Errorf("mkdirCustom: %w", err)
	}
	return nil
}

func (g *LayerGenerator) genLayerInterface() error {
	p := path.Join(g.targetLayerDirPath, g.tag, g.tag+".go")
	if err := templates.RunTemplate(&templates.TemplateData{
		Template: g.tmpl.Interface(),
		FilePath: p,
		Data:     g.data,
	}); err != nil {
		return fmt.Errorf("genLayerInterface: %w", err)
	}
	return nil
}

//func (g *LayerGenerator) genLayerDefaultImpl() error {
//	if err := g.genLayerInits(); err != nil {
//		return fmt.Errorf("genLayerDefaultImpl: %w", err)
//	}
//	if err := g.genLayerOperations(); err != nil {
//		return fmt.Errorf("genLayerDefaultImpl: %w", err)
//	}
//
//	return nil
//}

//func (g *LayerGenerator) genLayerInits() error {
//	if err := g.genLayerGeneratedInit(); err != nil {
//		return fmt.Errorf("genLayerInits: %w", err)
//	}
//	if err := g.genLayerCustomInit(); err != nil {
//		return fmt.Errorf("genLayerInits: %w", err)
//	}
//	return nil
//}

//func (g *LayerGenerator) genLayerGeneratedInit() error {
//	p := path.Join(g.getPathToGenerated(), g.tag+".go")
//	if err := g.runTemplate(p, g.tmpl.GeneratedInit()); err != nil {
//		return fmt.Errorf("genLayerGeneratedInit: %w", err)
//	}
//	return nil
//}

//func (g *LayerGenerator) genLayerCustomInit() error {
//	p := path.Join(g.getPathToCustom(), g.tag+".go")
//	if err := g.runTemplate(p, g.tmpl.CustomInit()); err != nil {
//		return fmt.Errorf("genLayerCustomInit: %w", err)
//	}
//	return nil
//}

func (g *LayerGenerator) getPathToDirBo() string {
	return path.Join(g.targetLayerDirPath, g.tag)
}

func (g *LayerGenerator) getPathToGenerated() string {
	return path.Join(g.targetLayerDirPath, g.tag, "generated")
}

func (g *LayerGenerator) getPathToCustom() string {
	return path.Join(g.targetLayerDirPath, g.tag, "custom")
}
