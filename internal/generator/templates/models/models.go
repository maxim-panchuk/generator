package models

import (
	"generator/internal/generator/utils"
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
		"upFirst":              utils.UpFirst,
		"lowFirst":             utils.LowFirst,
		"getField":             utils.GetField,
		"getDtoFieldType":      utils.GetDtoFieldType,
		"containsTime":         utils.ContainsTime,
		"getModelEnums":        utils.GetModelEnums,
		"toUpper":              strings.ToUpper,
		"getModelDependencies": utils.GetModelDependencies,
		"getRootFolderPath":    utils.GetRootFolderPath,
	}

	tmpl, err := template.New("models").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}
