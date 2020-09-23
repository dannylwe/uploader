package database

import(
	"github.com/danny/services/model"
	log "github.com/sirupsen/logrus"
)

func saveToDatabase(records [][]string) {
	log.Info("saving to database....")
	for _, record := range records[1:] {
		insertRecord := model.Sales{
			Region:record[0], Country:record[1], ItemType:record[2], SalesChannel:record[3], 
			OrderPriority:record[4], OrderDate:record[5], OrderID:record[6], ShipDate:record[7], 
			UnitsSold:record[8], UnitPrice:record[9], TotalRevenue:record[10], TotalCost:record[11], 
			TotalProfit:record[12]}
			
		result := model.DB.Create(&insertRecord)
		if result.Error != nil {
			log.Error(result.Error)
		}
	}
	log.Info("Completed saving recoreds")
	defer model.DB.Close()
	return
}