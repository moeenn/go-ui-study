package templates

import (
	"fmt"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func Load(path string) (*Templates, error) {
	tmpl, err := template.ParseGlob(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load templates: %w", err)
	}

	return &Templates{
		templates: tmpl,
	}, nil
}

func (t *Templates) Execute(writer io.Writer, templateName TemplateName, data any) error {
	return t.templates.ExecuteTemplate(writer, string(templateName), data)
}
