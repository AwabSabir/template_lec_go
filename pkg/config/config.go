package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	TamplateCashe map[string]*template.Template
	InfoLoger     *log.Logger
}
