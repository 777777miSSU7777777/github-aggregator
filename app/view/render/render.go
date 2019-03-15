package render

import (
	"html/template"
)

var templates *template.Template

func init(){
	templates = template.Must(template.ParseGlob("web/templates/*.gohtml"))
}

//SetTemplates set templates for rendering.
func SetTemplates(tpls *template.Template){
	templates = tpls
}

//GetTemplates get templates for rendering.
func GetTemplates()(*template.Template){
	return templates
}