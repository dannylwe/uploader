package common

import (
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// confirmFileType reads the file type 
func confirmFileType(file multipart.File, w http.ResponseWriter) error {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		RenderError(w, "INVALID_FILE\n", http.StatusBadRequest)
		return err
	}

	// check file type
	detectedFileType := http.DetectContentType(fileBytes)

	types, err := mime.ExtensionsByType(detectedFileType)
	if err != nil {
		RenderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
		return err
	}
	log.Info("File Type " + detectedFileType)
	log.Info(types)
	return nil
}
