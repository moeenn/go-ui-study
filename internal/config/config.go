package config

import (
	"fmt"
	"time"
)

type ServerConfig struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		Host:         "0.0.0.0",
		Port:         "8000",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 20,
	}
}

func (c ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

type TemplatesConfig struct {
	TemplatesPath string
}

func NewTemplatesConfig() TemplatesConfig {
	return TemplatesConfig{
		TemplatesPath: "./views/*.html",
	}
}

type Config struct {
	Server    ServerConfig
	Templates TemplatesConfig
}

func NewConfig() Config {
	server := NewServerConfig()
	templates := NewTemplatesConfig()

	return Config{
		Server:    server,
		Templates: templates,
	}
}
