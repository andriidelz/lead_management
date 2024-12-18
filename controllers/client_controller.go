// controllers/client_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"time"

	"lead_management/models"
	"lead_management/services"

	"github.com/gin-gonic/gin"
)

var clients = []models.Client{}
var nextID = 1

func CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client.ID = nextID
	nextID++
	client.LeadsAssigned = 0
	clients = append(clients, client)
	c.JSON(http.StatusCreated, client)
}

func GetAllClients(c *gin.Context) {
	c.JSON(http.StatusOK, clients)
}

func GetClientByID(c *gin.Context) {
	id := c.Param("id")
	for _, client := range clients {
		if id == strconv.Itoa(client.ID) {
			c.JSON(http.StatusOK, client)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
}

func AssignLead(c *gin.Context) {
	currentTime := time.Now().Format("15:04")
	bestClient := services.DetermineBestClient(clients, currentTime)
	if bestClient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No available clients"})
		return
	}
	bestClient.LeadsAssigned++
	c.JSON(http.StatusOK, gin.H{"client_assigned": bestClient.Name})
}
