package repository

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
var fInterface embed.FS

func (t *Template) Interface() *template.Template {
	f, err := fInterface.ReadFile("interface.tmpl")
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

	tmpl, err := template.New("repository interface").Funcs(funcMap).Parse(string(f))
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
		"lowFirst":             utils.LowFirst,
		"upFirst":              utils.UpFirst,
		"convertToGoType":      utils.ConvertToGoType,
		"getResponse":          utils.GetResponse,
		"getRootFolderPath":    utils.GetRootFolderPath,
		"getTagDependencies":   utils.GetTagDependencies,
		"getOperationCrudType": utils.GetOperationCrudType,
	}

	tmpl, err := template.New("repository generated").Funcs(funcMap).Parse(string(f))
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

	tmpl, err := template.New("repository custom").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}
