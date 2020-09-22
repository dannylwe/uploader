package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danny/services/common"
	"github.com/danny/services/model"
	log "github.com/sirupsen/logrus"
)

// GetProfitsByDate get profits by date range
func GetProfitsByDate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		log.Info("get profits by date range, limit 10")

		var date model.Dates
		const dbISOLAyout string = "2006-01-02"
		err := json.NewDecoder(r.Body).Decode(&date)

		from, _ := time.Parse(dbISOLAyout, date.StartDate)
		to, _ := time.Parse(dbISOLAyout, date.EndDate)

		if err != nil {
			log.Error(err)
			return
		}

		var profit model.Profit

		err = model.Db.QueryRow("SELECT SUM(total_profit) AS profit FROM sales WHERE order_date BETWEEN ? AND ?", from, to).Scan(&profit.Profit)
		if err != nil {
			log.Error(err)
			return
		}

		returnObject, _ := json.Marshal(profit)
		common.JsonResponse(w, returnObject)
		return

	}
	log.Info("Invalid HTTP method accessed")
	common.RenderError(w, "INVALID_METHOD", http.StatusMethodNotAllowed)
	return
}
