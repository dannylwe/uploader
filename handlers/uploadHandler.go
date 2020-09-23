package handlers

import (
	"net/http"

	"github.com/danny/services/common"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		common.RenderError(w, "INVALID METHOD", http.StatusBadRequest)
	case http.MethodPost:
		UploadFile(w, r)
	}
}
