package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafabcanedo/api-golang/src/configuration/logger"
	"github.com/rafabcanedo/api-golang/src/controller"
	"github.com/rafabcanedo/api-golang/src/controller/routes"
	"github.com/rafabcanedo/api-golang/src/model/service"
)

func main() {
	logger.Info("About to start the application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init mongodb
	

	// Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}