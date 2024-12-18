// services/lead_service.go
package services

import (
	"lead_management/models"
	"time"
)

func isWithinWorkingHours(currentTime string, workingHours models.TimeRange) bool {
	layout := "15:04"
	current, _ := time.Parse(layout, currentTime)
	start, _ := time.Parse(layout, workingHours.Start)
	end, _ := time.Parse(layout, workingHours.End)
	return current.After(start) && current.Before(end)
}

func DetermineBestClient(clients []models.Client, currentTime string) *models.Client {
	var bestClient *models.Client
	for i, client := range clients {
		if client.LeadsAssigned >= client.LeadCapacity || !isWithinWorkingHours(currentTime, client.WorkingHours) {
			continue
		}
		if bestClient == nil || client.Priority > bestClient.Priority {
			bestClient = &clients[i]
		}
	}
	return bestClient
}
