package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var dir, _ = os.Getwd()
var newDir = strings.Replace(dir, "\\", "/", -1)

var pathing, _ = filepath.Split(newDir)
var absPath, _ = filepath.Abs("./public/upload.html")

// for test
// var templates = template.Must(template.ParseFiles(pathing + "public/upload.html"))

// for application run
var templates = template.Must(template.ParseFiles(filepath.Join(newDir, "./public/upload.html")))


// Display template
func Display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}
