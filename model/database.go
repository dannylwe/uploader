package model

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	// need mysql connection
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB
var Db *sql.DB

// ConnectDatabase Connect to MYSQL database
func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Info("Successlly connected to database")

	// database.AutoMigrate(&Sales{})
	DB = database
	DB.Exec("CREATE TABLE IF NOT EXISTS sales(country VARCHAR(255) NOT NULL, region VARCHAR(255) NOT NULL,item_type VARCHAR(255) NOT NULL,sales_channel VARCHAR(255) NOT NULL,order_priority VARCHAR(255) NOT NULL,order_date DATE NOT NULL,order_id VARCHAR(255) NOT NULL,ship_date DATE NOT NULL,units_sold VARCHAR(255) NOT NULL,unit_price VARCHAR(255) NOT NULL,unit_cost VARCHAR(255) NOT NULL,total_revenue VARCHAR(255) NOT NULL,total_cost VARCHAR(255) NOT NULL,total_profit VARCHAR(255) NOT NULL)")

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(15 * time.Minute)
}

// SQLConn Create connection to dabase to run RAW query
func SQLConn() {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}
	Db = db
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	log.Info("SQL connection established")
}
