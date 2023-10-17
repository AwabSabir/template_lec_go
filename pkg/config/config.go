package config

import (
	"log"
	"text/template"
)

type AppConfig struct {
	UseCache      bool
	TamplateCashe map[string]*template.Template
	InfoLoger     *log.Logger
}
