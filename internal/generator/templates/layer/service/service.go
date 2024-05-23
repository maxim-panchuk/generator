package service

import (
	"embed"
	"generator/internal/generator/utils"
	"text/template"
)

type Template struct{}

func New() *Template {
	return &Template{}
}

//go:embed interface.tmpl
var fIntr embed.FS

func (t *Template) Interface() *template.Template {
	f, err := fIntr.ReadFile("interface.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"lowFirst":           utils.LowFirst,
		"upFirst":            utils.UpFirst,
		"convertToGoType":    utils.ConvertToGoType,
		"getResponse":        utils.GetResponse,
		"getRootFolderPath":  utils.GetRootFolderPath,
		"getTagDependencies": utils.GetTagDependencies,
	}

	tmpl, err := template.New("service interface").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}

//go:embed generated.tmpl
var fGenerated embed.FS

func (t *Template) GeneratedInit() *template.Template {
	f, err := fGenerated.ReadFile("generated.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"lowFirst":               utils.LowFirst,
		"upFirst":                utils.UpFirst,
		"convertToGoType":        utils.ConvertToGoType,
		"getResponse":            utils.GetResponse,
		"getRootFolderPath":      utils.GetRootFolderPath,
		"responseContainsSchema": utils.ResponseContainsSchema,
		"getTagDependencies":     utils.GetTagDependencies,
	}

	tmpl, err := template.New("service generated init").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}

//go:embed custom.tmpl
var fCustom embed.FS

func (t *Template) CustomInit() *template.Template {
	f, err := fCustom.ReadFile("custom.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"getRootFolderPath": utils.GetRootFolderPath,
	}

	tmpl, err := template.New("service custom init").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}
