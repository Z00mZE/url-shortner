package embedded

import (
	"embed"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

//go:embed template/*
var assetData embed.FS

// path имя папки c html шаблонами для встраивания в приложение
const path = "template/*.html"

// TemplateRenderer кастомный рендер `html/template`
type TemplateRenderer struct {
	templates *template.Template
}

// NewRenderer конструктор рендера html-шаблонов
func NewRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(
			template.ParseFS(
				assetData,
				path,
			),
		),
	}
}

// Render рендер html-шаблона
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
