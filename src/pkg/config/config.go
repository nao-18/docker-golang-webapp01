package config

import (
	"html/template"
)

// AppConfig holds the application configd
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
