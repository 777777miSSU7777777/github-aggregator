package view

import (
	"html/template"
)

var templates *template.Template


func init() {
	templates = template.Must(template.ParseGlob("../../web/templates/*.gohtml"))
}


func SetTemplates(tpls *template.Template) {
	templates = tpls
}


func GetTemplates() *template.Template {
	return templates
}
