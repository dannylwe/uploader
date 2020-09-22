package main

import (
	"github.com/danny/services/handlers"
	"github.com/danny/services/model"
)

func main() {
	model.ConnectDatabase()
	model.SQLConn()
	handlers.SetupRoutes()
}
