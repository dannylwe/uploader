package handlers

import (
	"github.com/danny/services/common"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		common.Display(w, "upload", nil)
	case http.MethodPost:
		UploadFile(w, r)
	}
}
