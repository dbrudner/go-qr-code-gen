package db

import (
	"log"
)

func createUserTable() {
	createUserTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE
);`

	_, err := DB.Exec(createUserTableSQL)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}

func createSiteTable() {
	createSiteTableSQL := `
    CREATE TABLE IF NOT EXISTS sites (
      id SERIAL PRIMARY KEY,
      name TEXT NOT NULL UNIQUE,
      url TEXT NOT NULL UNIQUE,
    );`

	_, err := DB.Exec(createSiteTableSQL)
	if err != nil {
		log.Fatalf("Failed to create sites table: %v", err)
	}
}

func createTicketTable() {
	createTicketTableSQL := `
    CREATE TABLE IF NOT EXISTS tickets (
      id SERIAL PRIMARY KEY,
      site_id INTEGER NOT NULL,
      FOREIGN KEY (site_id) REFERENCES sites(id),
      user_id INTEGER NOT NULL  ,
      FOREIGN KEY (user_id) REFERENCES users(id),
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      content TEXT NOT NULL,
    );`
	_, err := DB.Exec(createTicketTableSQL)
	if err != nil {
		log.Fatalf("Failed to create tickets table: %v", err)
	}
}

func CreateTables() {
	createUserTable()
	createSiteTable()
	createTicketTable()
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
