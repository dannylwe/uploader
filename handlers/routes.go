package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// SetupRoutes route handler for application
func SetupRoutes() {
	PORT := ":8080"
	log.Info("Starting application on port" + PORT)

	r := mux.NewRouter()
	r.Use(CORS)

	r.HandleFunc("/upload", UploadHandler)
	r.HandleFunc("/", RedirectToUpload)
	r.HandleFunc("/records", GetAllRecords)
	r.HandleFunc("/profit", GetProfitsByDate)
	r.HandleFunc("/topfive", GetTopFiveProfitableItems)
	http.ListenAndServe(PORT, r)
}

// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		log.Info("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

