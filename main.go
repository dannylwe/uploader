package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/danny/services/model"
)

func main() {
	model.ConnectDatabase()
	model.SQLConn()
	setupRoutes()
}

func setupRoutes() {
	PORT := ":8080"
	log.Info("Starting application on port" + PORT)

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", redirectToUpload)
	http.ListenAndServe(PORT, nil)
}

// compile template
var templates = template.Must(template.ParseFiles("public/upload.html"))

// Display template
func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		display(w, "upload", nil)
	case http.MethodPost:
		uploadFile(w, r)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	var maxUploadSize int64
	maxUploadSize = 15 * 1024000
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

	// TODO: check file type
	// err = confirmFileType(file, w)
	// if err != nil {
	// 	log.Warn(err)
	// 	return
	// }

	if err = validateFileSize(handler.Size, maxUploadSize, w); err != nil {
		log.Warn(err)
		return
	}

	defer file.Close()

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
	log.Info("Uploaded file " + handler.Filename)
	log.Info("Header size")
	log.Info(handler.Size)

	// read csv from disk
	records := readCSV(handler.Filename)
	// fmt.Println(records[1:])
	saveToDatabase(records)
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
	if fileSize > maxSize {
		renderError(w, "File Too Large", http.StatusRequestEntityTooLarge)
		return errors.New("File too big")
	}
	return nil
}

func confirmFileType(file multipart.File, w http.ResponseWriter) error {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		renderError(w, "INVALID_FILE\n", http.StatusBadRequest)
		return err
	}

	// check file type
	detectedFileType := http.DetectContentType(fileBytes)

	types, err := mime.ExtensionsByType(detectedFileType)
	if err != nil {
		renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
		return err
	}
	log.Info("File Type " + detectedFileType)
	log.Info(types)
	return nil
}

func readCSV(filename string) [][]string {
	log.Info("reading csv file " + filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// parse file
	r := csv.NewReader(file)
	r.Comma = ','

	rows, err := r.ReadAll()
	if err != nil {
		log.Error(err)
	}
	return rows
}

func saveToDatabase(records [][]string) {
	for _, record := range records[1:] {
		insert := model.Sales{record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12]}
		result := model.DB.Create(&insert)
		if result.Error != nil {
			log.Error(result.Error)
		}
	}
	log.Info("Completed saving recoreds")
	defer model.DB.Close()
	return
}