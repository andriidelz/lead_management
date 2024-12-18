// models/client.go
package models

type Client struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Priority      int       `json:"priority"`
	LeadCapacity  int       `json:"lead_capacity"`
	LeadsAssigned int       `json:"leads_assigned"`
	WorkingHours  TimeRange `json:"working_hours"`
}

type TimeRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
