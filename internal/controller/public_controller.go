package controller

import (
	"net/http"
	"ui/internal/templates"
)

func (c Controller) HomePage(w http.ResponseWriter, r *http.Request) {
	if err := c.Templates.Execute(w, templates.TemplateHomePage, nil); err != nil {
		c.Logger.Error("failed to execute home_page template", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
