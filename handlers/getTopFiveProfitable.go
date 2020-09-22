package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/danny/services/common"
	"github.com/danny/services/model"
	log "github.com/sirupsen/logrus"
)

// GetTopFiveProfitableItems get top five profitable items in period
func GetTopFiveProfitableItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var date model.Dates
		const dbISOLAyout string = "2006-01-02"
		err := json.NewDecoder(r.Body).Decode(&date)
		if err != nil {
			log.Error(err)
			return
		}

		fmt.Println(date)
		log.Info("get top five profitable items")

		var profit []model.TopProfitable

		from, _ := time.Parse(dbISOLAyout, date.StartDate)
		to, _ := time.Parse(dbISOLAyout, date.EndDate)

		rows, err := model.Db.Query("select item_type AS name, ROUND(SUM(total_profit), 2) AS profit from sales WHERE order_date BETWEEN ? AND ? GROUP BY item_type ORDER BY Profit DESC limit 5", from, to)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			var name string
			var profitable float64
			err = rows.Scan(&name, &profitable)
			if err != nil {
				log.Error(err)
			}
			total := model.TopProfitable{Name: name, Profit: profitable}
			profit = append(profit, total)
		}

		returnObject, err := json.Marshal(profit)
		if err != nil {
			fmt.Println(err)
		}
		common.JsonResponse(w, returnObject)
		return
	}
	log.Info("Invalid HTTP method accessed")
	common.RenderError(w, "INVALID_METHOD", http.StatusMethodNotAllowed)
}
