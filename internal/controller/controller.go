package controller

import (
	"log/slog"
	"net/http"
	"ui/internal/templates"
)

type Controller struct {
	Logger    *slog.Logger
	Templates *templates.Templates
}

func NewController(logger *slog.Logger, templates *templates.Templates) Controller {
	return Controller{
		Logger:    logger,
		Templates: templates,
	}
}

func (c Controller) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", c.HomePage)
	mux.HandleFunc("GET /login", c.LoginPage)
	mux.HandleFunc("GET /forgot-password", c.ForgotPassword)
}
