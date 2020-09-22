package handlers

import "net/http"

// RedirectToUpload redirects to upload URI
func RedirectToUpload(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}