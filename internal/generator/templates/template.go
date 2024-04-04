package templates

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"text/template"
)

type TemplateData struct {
	Template *template.Template
	FilePath string
	Data     any
}

//func RunTemplate(td *TemplateData) error {
//	var processed bytes.Buffer
//	if err := td.Template.Execute(&processed, td.Data); err != nil {
//		panic(err)
//	}
//
//	//formatted, err := format.Source(processed.Bytes())
//	//if err != nil {
//	//	return fmt.Errorf("execute template: %e", err)
//	//}
//
//	wrF, err := os.Create(td.FilePath)
//	if err != nil {
//		panic(err)
//	}
//	defer wrF.Close()
//
//	if _, err := wrF.Write(processed.Bytes()); err != nil {
//		return fmt.Errorf("execute template: %e", err)
//	}
//	return nil
//}

func RunTemplate(td *TemplateData) error {
	var processed bytes.Buffer
	if err := td.Template.Execute(&processed, td.Data); err != nil {
		panic(err)
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		return fmt.Errorf("execute template: %e", err)
	}

	wrF, err := os.Create(td.FilePath)
	if err != nil {
		panic(err)
	}
	defer wrF.Close()

	if _, err := wrF.Write(formatted); err != nil {
		return fmt.Errorf("execute template: %e", err)
	}
	return nil
}
