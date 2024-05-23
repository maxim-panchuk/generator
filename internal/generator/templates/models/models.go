package models

import (
	"embed"
	_ "embed"
	"generator/internal/generator/utils"
	"strings"
	"text/template"
)

//go:embed models.tmpl
var modelsF embed.FS

func Models() *template.Template {
	f, err := modelsF.ReadFile("models.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"upFirst":                     utils.UpFirst,
		"lowFirst":                    utils.LowFirst,
		"getField":                    utils.GetField,
		"getDtoFieldType":             utils.GetDtoFieldType,
		"containsTime":                utils.ContainsTime,
		"getModelEnums":               utils.GetModelEnums,
		"toUpper":                     strings.ToUpper,
		"getModelDependencies":        utils.GetModelDependencies,
		"getRootFolderPath":           utils.GetRootFolderPath,
		"isModelEntity":               utils.IsModelEntity,
		"getModelPrimaryKeyField":     utils.GetModelPrimaryKeyField,
		"getTypeForEntity":            utils.GetTypeForEntity,
		"getAnnotationForEntityField": utils.GetAnnotationForEntityField,
	}

	tmpl, err := template.New("models").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}
