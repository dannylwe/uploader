package common

import (
	"html/template"
	"net/http"
)

// compile template
var templates = template.Must(template.ParseFiles("public/upload.html"))

// Display template
func Display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}