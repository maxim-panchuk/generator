package mapper

import (
	"embed"
	_ "embed"
	"generator/internal/generator/utils"
	"text/template"
)

//go:embed mapper.tmpl
var f embed.FS

func Mapper() *template.Template {
	f, err := f.ReadFile("mapper.tmpl")
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
