package db

import (
	"fmt"
	"log"
	"os"
)

func executeSQLFromFile(filename string) {
	// Read SQL file
	sqlBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	// Convert to string
	sql := string(sqlBytes)

	fmt.Println("sql:", sql)
	// Execute SQL
	_, err = DB.Exec(sql)
	if err != nil {
		log.Fatalf("Failed to execute SQL: %v", err)
	}
}

func createUsersTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}

func createTicketsTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS tickets (id TEXT PRIMARY KEY, site_id INTEGER NOT NULL, user_id INTEGER NOT NULL, created_at TIMESTAMP DEFAULT current_timestamp, updated_at TIMESTAMP DEFAULT current_timestamp, content TEXT NOT NULL)")
	if err != nil {
		log.Fatalf("Failed to create tickets table: %v", err)
	}
}

func createSitesTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS sites (id TEXT PRIMARY KEY, description TEXT NOT NULL, url TEXT NOT NULL UNIQUE)")
	if err != nil {
		log.Fatalf("Failed to create sites table: %v", err)
	}
}

func CreateTables() {
	createUsersTable()
	createSitesTable()
	createTicketsTable()
}

type User = struct {
	name  string
	email string
}

func SeedData() {
	users := []User{
		{"Dave Brudner", "davebrudner@gmail.com"},
	}

	for _, user := range users {
		_, err := DB.Exec("INSERT INTO users (email) VALUES (?)", user.email)
		if err != nil {
			log.Printf("Failed to seed user %s: %v", user.email, err)
		}
	}

	log.Println("Seeding completed")
}
