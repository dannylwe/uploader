package model

// import "github.com/jinzhu/gorm"

// Sales is struct to database
type Sales struct {
	Region       string
	Country      string
	ItemType     string
	SalesChannel string
	OrderPrice   string
	OrderDate    string
	OrderID      string
	ShipDate     string
	UnitsSold    string
	UnitPrice    string
	TotalRevenue string
	TotalCost    string
	TotalProfit  string
}
