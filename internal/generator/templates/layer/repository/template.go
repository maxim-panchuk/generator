package repository

import "text/template"

type Template struct{}

func New() *Template {
	return &Template{}
}

func (t *Template) Interface() *template.Template {
	return nil
}

func (t *Template) GeneratedImpl() *template.Template {
	return nil
}

func (t *Template) GeneratedInit() *template.Template {
	return nil
}

func (t *Template) CustomInit() *template.Template {
	return nil
}
