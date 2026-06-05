package main

import (
	"log"

	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/controller"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/controller/routes"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init depedencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
