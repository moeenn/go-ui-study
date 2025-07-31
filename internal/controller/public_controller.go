package controller

import (
	"net/http"
	"ui/internal/templates"
)

func (c Controller) HomePage(w http.ResponseWriter, r *http.Request) {
	args := templates.HomePageArgs{
		Headings: []templates.HeadingSection{
			{Level: 1, Heading: "Heading one", Description: "Some random description"},
			{Level: 3, Heading: "Heading of three", Description: "Some heading three random description"},
			{Level: 5, Heading: "Another heading", Description: "Another simple description"},
		},
	}

	if err := c.Templates.Execute(w, templates.TemplateHomePage, args); err != nil {
		c.Logger.Error("failed to execute home_page template", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
