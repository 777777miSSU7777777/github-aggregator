package view

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetTemplate__SameTemplate__Equal(t *testing.T) {
	templ := &template.Template{}

	SetTemplates(templ)

	assert.Equal(t, templ, templates)
}

func TestGetTemplate__SameTemplate__Equal(t *testing.T) {
	templ := &template.Template{}

	SetTemplates(templ)

	assert.Equal(t, templ, GetTemplates())
}
