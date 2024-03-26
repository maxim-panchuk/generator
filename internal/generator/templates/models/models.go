package models

import (
	"generator/internal/generator/utils"
	"os"
	"text/template"
)

func Models() *template.Template {
	f, err := os.ReadFile("internal/generator/templates/models/models.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"upFirst":         utils.UpFirst,
		"getField":        utils.GetField,
		"getDtoFieldType": utils.GetDtoFieldType,
	}

	tmpl, err := template.New("models").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}
