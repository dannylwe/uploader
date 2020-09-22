package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danny/services/model"
	log "github.com/sirupsen/logrus"
)

// GetAllRecords gets all records in descending order. Limited to 10
func GetAllRecords(w http.ResponseWriter, r *http.Request) {
	log.Info("Getting all records, limit 10")
	limit := 10
	var sales []model.Sales
	if err := model.DB.Order("order_date desc").Limit(limit).Find(&sales).Error; err != nil {
		log.Error(err)
		return
	}
	log.Info("get all records limit 10 SUCCESS")
	returnObject, _ := json.Marshal(sales)
	w.Header().Set("Content-Type", "application/json")
  	w.Write(returnObject)
	return
}