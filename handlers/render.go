package handlers

import (
	"html/template"
	"net/http"
	"os"
)
var wd, _ = os.Getwd()

// compile template
var templates = template.Must(template.ParseFiles("C:/Users/pc/Desktop/projects/uploader/"+ "/public/upload.html"))

// Display template
func Display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}
