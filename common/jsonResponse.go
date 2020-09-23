package common

import "net/http"

// JsonResponse responds with JSON
func JsonResponse(w http.ResponseWriter, ResponseObject []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(ResponseObject)
	return
}
