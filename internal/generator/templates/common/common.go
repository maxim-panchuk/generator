package common

import (
	"embed"
	"generator/internal/generator/utils"
	"strings"
	"text/template"
)

//go:embed main.tmpl
var fM embed.FS

func Main() *template.Template {
	f, err := fM.ReadFile("main.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"lowFirst":          utils.LowFirst,
		"getServiceName":    utils.GetServiceName,
		"upFirst":           utils.UpFirst,
		"convertToGoType":   utils.ConvertToGoType,
		"getRootFolderPath": utils.GetRootFolderPath,
		"toUpper":           strings.ToUpper,
	}

	tmpl, err := template.New("controller interface").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}
