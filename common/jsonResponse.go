package common

import "net/http"

func JsonResponse(w http.ResponseWriter, ResponseObject []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(ResponseObject)
	return
}