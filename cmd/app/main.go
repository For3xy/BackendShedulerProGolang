package main

import (
	"backendShedulerPro/internal/database"
	"backendShedulerPro/internal/database/controllers"
	"backendShedulerPro/internal/models"
	"backendShedulerPro/internal/transport/rest/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	MailingController       controllers.MailingController
	MailingRoutesController routes.MailingRoutesController
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Mailing{})

	MailingController = controllers.NewMailingController(database.DB)
	MailingRoutesController = routes.NewMailingController(MailingController)
	server = gin.Default()

	server.Use(cors.Default())

	router := server.Group("/api")
	router.GET("/checker", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "success", "data": "Checker API is working!"})
	})

	MailingRoutesController.MailingRoute(router)
	log.Fatal(server.Run(":8080"))
}
