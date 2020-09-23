package config

import (
	
	"github.com/danny/services/handlers"
	"github.com/danny/services/model"
)

// Run initializes and starts the application
func Run() {
	model.ConnectDatabase()
	model.SQLConn()
	handlers.SetupRoutes()
}
