// Package view parses html templates for web version of github aggregator.
// Also contains getter and setter for templates.
package view

import (
	"html/template"
)

var templates *template.Template

// SetTemplates sets templates for rendering.
func SetTemplates(tpls *template.Template) {
	templates = tpls
}

// GetTemplates gets templates for rendering.
func GetTemplates() *template.Template {
	return templates
}
