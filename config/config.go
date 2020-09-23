package config

import (
	"os"
	"runtime"
	"strings"

	"github.com/danny/services/handlers"
	"github.com/danny/services/model"
)

// Run initializes and starts the application
func Run() {
	model.ConnectDatabase()
	model.SQLConn()
	handlers.SetupRoutes()
}

func GetWD() string{
	dir, _ := os.Getwd()
	var ss [] string
  	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
		 
	}
	return ss[len(ss)-1]
}