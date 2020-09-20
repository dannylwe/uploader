package main

import (
	"fmt"
	"net/http"
	"text/template"
)


func main() {
	http.HandleFunc("/upload", uploadHandler)

	http.ListenAndServe(":9000", nil)
}

// compile template
var templates = template.Must(template.ParseFiles("public/upload.html"))

// Display template
func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page + ".html", data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "upload", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File")
}