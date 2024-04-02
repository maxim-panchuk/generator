package controller

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
	f, err := os.ReadFile("internal/generator/templates/layer/controller/interface.tmpl")
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

func (t *Template) GeneratedInit() *template.Template {
	f, err := os.ReadFile("internal/generator/templates/layer/controller/generated.tmpl")
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

func (t *Template) CustomInit() *template.Template {
	f, err := os.ReadFile("internal/generator/templates/layer/controller/custom.tmpl")
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
