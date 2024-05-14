package mapper

import (
	"generator/internal/generator/utils"
	"os"
	"text/template"
)

func Mapper() *template.Template {
	f, err := os.ReadFile("internal/generator/templates/mapper/mapper.tmpl")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"lowFirst":             utils.LowFirst,
		"upFirst":              utils.UpFirst,
		"containsTime":         utils.ContainsTime,
		"getModelDependencies": utils.GetModelDependencies,
		"getRootFolderPath":    utils.GetRootFolderPath,
		"getField":             utils.GetField,
	}

	tmpl, err := template.New("mapper").Funcs(funcMap).Parse(string(f))
	if err != nil {
		panic(err)
	}

	return tmpl
}
