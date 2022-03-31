package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	Logger         *log.Logger
	IsProd         bool
	SessionManager *scs.SessionManager
}
