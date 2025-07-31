package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guiziin227/CRUDgo/src/configuration/logger"
	"github.com/guiziin227/CRUDgo/src/controller"
	"github.com/guiziin227/CRUDgo/src/controller/routes"
	"github.com/guiziin227/CRUDgo/src/model/service"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("Starting CRUDgo application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init dependencies
	service := service.NewUserDomainService()
	controller := controller.NewUserController(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
