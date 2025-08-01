package controller

import (
	"net/http"
	"ui/internal/templates"
)

func (c Controller) LoginPage(w http.ResponseWriter, r *http.Request) {
	if err := c.Templates.Execute(w, templates.TemplateLoginPage, nil); err != nil {
		c.Logger.Error("failed to render login page", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c Controller) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	if err := c.Templates.Execute(w, templates.TemplateForgotPasswordPage, nil); err != nil {
		c.Logger.Error("failed to render forgot-password page", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
