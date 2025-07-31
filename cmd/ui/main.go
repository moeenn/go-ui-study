package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"ui/internal/config"
	"ui/internal/controller"
	"ui/internal/templates"
)

func run() error {
	globalConfig := config.NewConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	templates, err := templates.Load(globalConfig.Templates.TemplatesPath)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public", fs))

	controller := controller.NewController(logger, templates)
	controller.RegisterRoutes(mux)

	address := globalConfig.Server.Address()
	logger.Info("starting server", "address", address)

	//nolint:exhaustruct
	server := http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  globalConfig.Server.ReadTimeout,
		WriteTimeout: globalConfig.Server.WriteTimeout,
	}

	return server.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
