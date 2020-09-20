package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Info("Successlly connected to database")

	database.AutoMigrate(&Sales{})
	DB = database
}
