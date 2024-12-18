// main.go
package main

import (
	"lead_management/database"
	"lead_management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	defer database.Close()
	database.InitializeTables()

	// router := gin.Default()
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.RegisterClientRoutes(router)

	router.Run(":8080")
}
