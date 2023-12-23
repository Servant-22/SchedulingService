package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// connectDB initialiseert en retourneert een databaseverbinding.
func connectDB() *sql.DB {
	// Haal databaseconfiguratie op uit omgevingsvariabelen
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Maak de database connectie string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Maak verbinding met de database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Fout bij het verbinden met de database: %v", err)
	}

	// Test de databaseverbinding
	err = db.Ping()
	if err != nil {
		log.Fatalf("Fout bij het pingen van de database: %v", err)
	}

	fmt.Println("Succesvol verbonden met de database")
	return db
}
