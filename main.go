package main

import (
	"20241126/config"
	"20241126/controller"
	"20241126/database"
	"20241126/middleware"
	"20241126/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.LoadConfig()
}

func main() {
	r := gin.Default()

	// Middleware
	r.Use(middleware.Logger())
	//r.Use(middleware.BasicAuth())

	controller := controller.NewController(database.GetDB())
	router.APIRouter(r, controller)

	r.Run(viper.GetString("SERVER_PORT"))
}
