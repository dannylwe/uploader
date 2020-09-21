package model

import "github.com/jinzhu/gorm"

// Sales is struct to database
type Sales struct {
	gorm.Model
	Region       string `json:"region"`
	Country      string	`json:"country"`
	ItemType     string `json:"ItemType"`
	SalesChannel string `json:"SalesChannel"`
	OrderPrice   string `json:"OrderPrice"`
	OrderDate    string `json:"OrderDate"`
	OrderID      string `json:"OrderID"`
	ShipDate     string `json:"ShipDate"`
	UnitsSold    string `json:"UnitsSold"`
	UnitPrice    string `json:"UnitPrice"`
	TotalRevenue string `json:"TotalRevenue"`
	TotalCost    string `json:"TotalCost"`
	TotalProfit  string `json:"TotalProfit"`
}
