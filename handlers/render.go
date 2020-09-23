package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// var wd = getWD()
var dir, _ = os.Getwd()
var newDir = strings.Replace(dir, "\\", "/", -1)
// var newer = strings.Split(newDir, "/")
var pathing, _ = filepath.Split(newDir)
var absPath, _ = filepath.Abs("./public/upload.html")
// compile template
// var templates = template.Must(template.ParseFiles(newDir + "/public/upload.html"))
// var templates = template.Must(template.ParseFiles(pathing + "public/upload.html"))

var templates = template.Must(template.ParseFiles(filepath.Join(newDir, "./public/upload.html")))


// Display template
func Display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

func getWD() string{
	dir, _ := os.Getwd()
	var ss [] string
  	if runtime.GOOS == "windows" {
		var st = strings.Replace(dir, "\\", "/", -1)
		ss = strings.Split(st, "/")
	} else {
		ss = strings.Split(dir, "/")
	}
	return ss[len(ss)-1]
}
