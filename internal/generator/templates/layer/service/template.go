package service

import (
	"generator/internal/generator/utils"
	"os"
	"text/template"
)

type Template struct{}

func New() *Template {
	return &Template{}
}

func (t *Template) Interface() *template.Template {
	f, err := os.ReadFile("internal/generator/templates/layer/service/interface.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"lowFirst":          utils.LowFirst,
		"upFirst":           utils.UpFirst,
		"convertToGoType":   utils.ConvertToGoType,
		"getResponse":       utils.GetResponse,
		"getRootFolderPath": utils.GetRootFolderPath,
	}

	tmpl, err := template.New("service interface").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}

func (t *Template) GeneratedInit() *template.Template {
	f, err := os.ReadFile("internal/generator/templates/layer/service/generated.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"lowFirst":          utils.LowFirst,
		"upFirst":           utils.UpFirst,
		"convertToGoType":   utils.ConvertToGoType,
		"getResponse":       utils.GetResponse,
		"getRootFolderPath": utils.GetRootFolderPath,
	}

	tmpl, err := template.New("service generated init").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}

func (t *Template) CustomInit() *template.Template {
	f, err := os.ReadFile("internal/generator/templates/layer/service/custom.tmpl")
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
