package common

import (
	"generator/internal/generator/utils"
	"os"
	"text/template"
)

func Main() *template.Template {
	f, err := os.ReadFile("internal/generator/templates/common/main.tmpl")
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
