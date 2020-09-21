package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"text/template"

	"github.com/danny/services/model"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
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
	http.HandleFunc("/records", getAllRecords)
	http.HandleFunc("/profit", getProfitsByDate)
	http.HandleFunc("/topfive", getTopFiveProfitableItems)
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

	filePath := "./" + handler.Filename
	mysql.RegisterLocalFile(filePath)
	err = model.DB.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' REPLACE INTO TABLE sales FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 LINES").Error
	if err != nil {
		log.Error("Could not load csv file to database")
	}
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

func getAllRecords(w http.ResponseWriter, r *http.Request) {
	log.Info("Getting all records, limit 10")
	limit := 10
	var sales []model.Sales
	if err := model.DB.Order("order_date desc").Limit(limit).Find(&sales).Error; err != nil {
		log.Error(err)
		return
	}
	log.Info("get all records limit 10 SUCCESS")
	returnObject, _ := json.Marshal(sales)
	w.Header().Set("Content-Type", "application/json")
  	w.Write(returnObject)
	return
}

func getProfitsByDate(w http.ResponseWriter, r *http.Request) {
	log.Info("get profits by date range, limit 10")

	var profit model.Profit
	model.DB.Raw("SELECT SUM(total_profit) AS profit FROM sales WHERE DATE(order_date) BETWEEN DATE(2016-09-09) AND DATE(2016-10-19)").Scan(&profit)
	
	returnObject, _ := json.Marshal(profit)
	w.Header().Set("Content-Type", "application/json")
	w.Write(returnObject)
	return
}

func getTopFiveProfitableItems(w http.ResponseWriter, r *http.Request) {
	log.Info("get top five profitable items")

	var profit []model.TopProfitable
	model.DB.Raw("select item_type AS name, ROUND(sum(total_profit), 2) AS profit from sales WHERE DATE(order_date) BETWEEN DATE(2016-09-09) AND DATE(2016-10-19) GROUP BY item_type ORDER BY Profit DESC limit 5").Scan(&profit)

	returnObject, _ := json.Marshal(profit)
	w.Header().Set("Content-Type", "application/json")
	w.Write(returnObject)
	return
}
