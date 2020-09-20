package model

type sales struct {
	Region       string
	Country      string
	ItemType     string
	SalesChannel string
	OrderPrice   string
	OrderDate    string
	OrderID      int64
	ShipDate     string
	UnitsSold    int32
	UnitPrice    float32
	TotalRevenue float64
	TotalCost    float64
	TotalProfit  float64
}
