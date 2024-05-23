package controller

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
		"upFirst":           utils.UpFirst,
		"convertToGoType":   utils.ConvertToGoType,
		"getRootFolderPath": utils.GetRootFolderPath,
	}

	tmpl, err := template.New("controller interface").Funcs(funcMap).Parse(string(f))
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
		"upFirst":                      utils.UpFirst,
		"lowFirst":                     utils.LowFirst,
		"convertToGoType":              utils.ConvertToGoType,
		"getRootFolderPath":            utils.GetRootFolderPath,
		"containsPathParameters":       utils.ContainsPathParameters,
		"getResponseByCode":            utils.GetResponseByCode,
		"responseContainsSchema":       utils.ResponseContainsSchema,
		"tagContainsTypicalOperation":  utils.TagContainsTypicalOperation,
		"getControllerTagDependencies": utils.GetControllerTagDependencies,
	}

	tmpl, err := template.New("controller interface").Funcs(funcMap).Parse(string(f))
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
		"upFirst":           utils.UpFirst,
		"convertToGoType":   utils.ConvertToGoType,
		"getRootFolderPath": utils.GetRootFolderPath,
	}

	tmpl, err := template.New("controller interface").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}
