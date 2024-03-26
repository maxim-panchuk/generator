package models

import (
	"os"
	"strings"
	"text/template"
)

func Models() *template.Template {
	f, err := os.ReadFile("internal/generator/templates/models/models.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"upFirst": strings.ToUpper,
	}

	tmpl, err := template.New("models").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}
