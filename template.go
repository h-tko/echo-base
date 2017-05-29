package main

import (
	"github.com/h-tko/echo-base/helpers"
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		templates: template.Must(template.New("").Funcs(helpers.TemplateHelpers).ParseGlob("app/views/**/*.html")),
	}
}
