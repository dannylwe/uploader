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
	"time"

	"github.com/danny/services/common"
	"github.com/danny/services/model"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func init() {
}

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
	err = model.DB.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' REPLACE INTO TABLE sales FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 LINES (region, country, item_type, sales_channel,order_priority, @order_date,order_id ,ship_date,units_sold, unit_price, unit_cost, total_revenue, total_cost, total_profit) SET order_date = STR_TO_DATE(@order_date, '%m/%d/%Y')").Error
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
	if r.Method == http.MethodPost {
		log.Info("get profits by date range, limit 10")

		var date model.Dates
		const dbISOLAyout string= "2006-01-02"
		err := json.NewDecoder(r.Body).Decode(&date)

		from, _ := time.Parse(dbISOLAyout, date.StartDate)
		to, _ := time.Parse(dbISOLAyout, date.EndDate)

		if err != nil {
			log.Error(err)
			return
		}
	
		var profit model.Profit
		
		err = model.Db.QueryRow("SELECT SUM(total_profit) AS profit FROM sales WHERE order_date BETWEEN ? AND ?", from, to).Scan(&profit.Profit)
		if err != nil {
			log.Error(err)
			return
		}

		returnObject, _ := json.Marshal(profit)
		common.JsonResponse(w, returnObject)
		return
		
	}
	log.Info("Invalid HTTP method accessed")
	renderError(w, "INVALID_METHOD", http.StatusMethodNotAllowed)
	return
}

func getTopFiveProfitableItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		
		var date model.Dates
		const dbISOLAyout string= "2006-01-02"
		err := json.NewDecoder(r.Body).Decode(&date)
		if err != nil {
			log.Error(err)
			return
		}

		fmt.Println(date)
		log.Info("get top five profitable items")

		var profit []model.TopProfitable

		from, _ := time.Parse(dbISOLAyout, date.StartDate)
		to, _ := time.Parse(dbISOLAyout, date.EndDate)
	
		rows, err := model.Db.Query("select item_type AS name, ROUND(SUM(total_profit), 2) AS profit from sales WHERE order_date BETWEEN ? AND ? GROUP BY item_type ORDER BY Profit DESC limit 5", from, to)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			var name string
			var profitable float64
			err = rows.Scan(&name, &profitable)
			if err != nil {
				log.Error(err)
			}
			total := model.TopProfitable{Name:name, Profit:profitable}
			profit = append(profit, total)
		}

		returnObject, err := json.Marshal(profit)
		if err != nil {
			fmt.Println(err)
		}
		common.JsonResponse(w, returnObject)
		return
	}
	log.Info("Invalid HTTP method accessed")
	renderError(w, "INVALID_METHOD", http.StatusMethodNotAllowed)
}
