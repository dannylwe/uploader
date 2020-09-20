package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)


func main() {
	setupRoutes()
}

func setupRoutes() {
	PORT := ":8080"
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
	var maxUploadSize int64
	maxUploadSize = 7 * 1024
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		fmt.Printf("Could not parse multipart form: %v\n", err)
    	renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
    	return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Printf("Error. Cannot get File %v\n", err)
		return
	}

	if err := validateFileSize(handler.Size, maxUploadSize, w); err != nil {
		fmt.Printf("%v\n", err)
		return;
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %v\n", handler.Filename)
	fmt.Printf("File Size: %v\n", handler.Size)
	fmt.Printf("MIME Header: %v\n", handler.Header)

	// create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		renderError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// save file to disk
	if _, err := io.Copy(dst, file); err != nil {
		renderError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded file")
	return
}

func redirectToUpload(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func validateFileSize(fileSize, maxSize int64, w http.ResponseWriter) error {
	if(fileSize > maxSize) {
		renderError(w, "File Too Large", http.StatusRequestEntityTooLarge)
		return errors.New("File too big")
	}
	return nil
}
