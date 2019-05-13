// Package view parses html templates for web version of github aggregator.
// Also contains getter and setter for templates.
package view

import (
	"html/template"

	"github.com/go-kit/kit/log"
)

var templates *template.Template

var logger log.Logger

func SetLogger(newLogger log.Logger) {
	logger = newLogger
}

func Logger() log.Logger {
	return logger
}

// AuthState struct which represents auth state of user.
type AuthState struct {
	Auth bool
}

// SetTemplates sets templates for rendering.
func SetTemplates(tpls *template.Template) {
	templates = tpls
}

// GetTemplates gets templates for rendering.
func GetTemplates() *template.Template {
	return templates
}
