package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Logger        *log.Logger
	IsProd        bool
}
