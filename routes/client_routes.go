// routes/client_routes.go
package routes

import (
	"lead_management/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterClientRoutes(router *gin.Engine) {
	clientGroup := router.Group("/clients")
	{
		clientGroup.POST("/", controllers.CreateClient)
		clientGroup.GET("/", controllers.GetAllClients)
		clientGroup.GET("/:id", controllers.GetClientByID)
		clientGroup.POST("/assign", controllers.AssignLead)
	}
}
