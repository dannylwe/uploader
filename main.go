package main

import (
	"fmt"
	"net/http"
	"text/template"
)


func main() {
	setupRoutes()
}

func setupRoutes() {
	PORT := ":9000"
	fmt.Println("Starting application on port" + PORT)
	
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", redirectToUpload)
	http.ListenAndServe(PORT, nil)
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
	maxUploadSize := 2 * 1024
	if err := r.ParseMultipartForm(int64(maxUploadSize)); err != nil {
		fmt.Printf("Could not parse multipart form: %v\n", err)
    	renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
    	return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Printf("Error. Cannot get File %v\n", err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %v", handler.Filename)
	fmt.Printf("File Size: %v", handler.Size)
	fmt.Printf("MIME Header: %v\n", handler.Header)
}

func redirectToUpload(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
