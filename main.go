package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"text/template"

	"github.com/danny/services/common"
	"github.com/danny/services/handlers"
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
	http.HandleFunc("/", handlers.RedirectToUpload)
	http.HandleFunc("/records", handlers.GetAllRecords)
	http.HandleFunc("/profit", handlers.GetProfitsByDate)
	http.HandleFunc("/topfive", handlers.GetTopFiveProfitableItems)
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
		common.RenderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
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

	if err = common.ValidateFileSize(handler.Size, maxUploadSize, w); err != nil {
		log.Warn(err)
		return
	}

	defer file.Close()

	// create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		common.RenderError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// save file to disk
	if _, err := io.Copy(dst, file); err != nil {
		common.RenderError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded file")
	log.Info("Uploaded file " + handler.Filename)
	log.Info("Header size")
	log.Info(handler.Size)

	filePath := "./" + handler.Filename
	mysql.RegisterLocalFile(filePath)
	err = model.DB.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' REPLACE INTO TABLE sales FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 LINES (region, country, item_type, sales_channel,order_priority, @order_date,order_id ,ship_date,units_sold, unit_price, unit_cost, total_revenue, total_cost, total_profit) SET order_date = STR_TO_DATE(@order_date, '%m/%d/%Y')").Error
	if err != nil {
		log.Error("Could not load csv file to database")
	}
	return
}

func confirmFileType(file multipart.File, w http.ResponseWriter) error {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		common.RenderError(w, "INVALID_FILE\n", http.StatusBadRequest)
		return err
	}

	// check file type
	detectedFileType := http.DetectContentType(fileBytes)

	types, err := mime.ExtensionsByType(detectedFileType)
	if err != nil {
		common.RenderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
		return err
	}
	log.Info("File Type " + detectedFileType)
	log.Info(types)
	return nil
}
