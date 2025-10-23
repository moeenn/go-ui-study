package controller

import (
	"net/http"
	"ui/internal/templates"
)

func (c *Controller) ErrorPage(w http.ResponseWriter, r *http.Request) {
	args := templates.ErrorPageArgs{
		StatusCode: 400,
		Error:      "You dont have access to the requested resource.",
	}

	if err := c.Templates.Execute(w, templates.TemplateErrorPage, args); err != nil {
		c.Logger.Error("failed to render error page", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
