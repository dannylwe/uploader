package common

import (
	"errors"
	"net/http")

// ValidateFileSize validates the size of the file
func ValidateFileSize(fileSize, maxSize int64, w http.ResponseWriter) error {
	if fileSize > maxSize {
		RenderError(w, "File Too Large", http.StatusRequestEntityTooLarge)
		return errors.New("File too big")
	}
	return nil
}