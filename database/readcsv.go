package database

import (
	"os"
	"encoding/csv"
	log "github.com/sirupsen/logrus"
)

func readCSV(filename string) [][]string {
	log.Info("reading csv file " + filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// parse file
	r := csv.NewReader(file)
	r.Comma = ','

	rows, err := r.ReadAll()
	if err != nil {
		log.Error(err)
	}
	return rows
}