package render

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(path string) *TemplateRenderer {
	return &TemplateRenderer{templates: template.Must(template.ParseGlob(path))}
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
