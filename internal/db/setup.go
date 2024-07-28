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

func CreateTables() {
	executeSQLFromFile("internal/db/queries/create_users_table.sql")
	executeSQLFromFile("internal/db/queries/create_sites_table.sql")
	executeSQLFromFile("internal/db/queries/create_tickets_table.sql")
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
