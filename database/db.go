package database

import (
	"context"
	"fmt"
	"lead_management/models"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var dbPool *pgxpool.Pool

func Connect() {
	var err error
	connString := "postgres://myuser:mypassword@localhost:5432/lead_management_db"
	dbPool, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Successfully connected to the database!")
}

func GetDB() *pgxpool.Pool {
	return dbPool
}

func Close() {
	dbPool.Close()
}

func InitializeTables() {
	conn := GetDB()

	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS clients (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100),
			priority INT,
			lead_capacity INT,
			leads_assigned INT,
			working_hours_start TIME,
			working_hours_end TIME
		);
	`)

	if err != nil {
		log.Fatalf("Unable to create tables: %v\n", err)
	}
}

func InsertClient(client models.Client) error {
	conn := GetDB()

	_, err := conn.Exec(context.Background(), `
		INSERT INTO clients (name, priority, lead_capacity, leads_assigned, working_hours_start, working_hours_end) 
		VALUES ($1, $2, $3, $4, $5, $6);
	`, client.Name, client.Priority, client.LeadCapacity, client.LeadsAssigned, client.WorkingHours.Start, client.WorkingHours.End)

	return err
}

func GetAllClients() ([]models.Client, error) {
	conn := GetDB()
	rows, err := conn.Query(context.Background(), "SELECT id, name, priority, lead_capacity, leads_assigned, working_hours_start, working_hours_end FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Priority, &client.LeadCapacity, &client.LeadsAssigned, &client.WorkingHours.Start, &client.WorkingHours.End); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}
