package handlers

import (
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		Display(w, "upload", nil)
	case http.MethodPost:
		UploadFile(w, r)
	}
}
