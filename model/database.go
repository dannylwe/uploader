package model

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
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

	database.AutoMigrate(&Sales{})
	DB = database
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