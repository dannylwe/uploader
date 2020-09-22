package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// SetupRoutes route handler for application
func SetupRoutes() {
	PORT := ":8080"
	log.Info("Starting application on port" + PORT)

	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/", RedirectToUpload)
	http.HandleFunc("/records", GetAllRecords)
	http.HandleFunc("/profit", GetProfitsByDate)
	http.HandleFunc("/topfive", GetTopFiveProfitableItems)
	http.ListenAndServe(PORT, nil)
}
