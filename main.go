package main

import (
	"basicmigrate/db"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define flags for migrations
	migrateUp := flag.Bool("up", false, "Apply all migrations")
	migrateDown := flag.Bool("down", false, "Rollback all migrations")
	specificDown := flag.String("undo", "", "Undo specific migration by ID (e.g., '0001_create_users')")
	flag.Parse()

	// Connect to the database
	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Check for migration flags
	if *migrateUp {
		fmt.Println("Running migrations...")
		db.RunMigrations(dbConn)
		return // Exit after applying migrations
	} else if *migrateDown {
		fmt.Println("Rolling back all migrations...")
		db.RollbackMigrations(dbConn)
		return // Exit after rolling back migrations
	} else if *specificDown != "" {
		fmt.Printf("Rolling back migration: %s\n", *specificDown)
		db.RollbackMigration(dbConn, *specificDown)
		return // Exit after rolling back the specific migration
	}

	// Start the web server if no migration flags are set
	fmt.Println("Starting server on port 8080...")

	// Simple HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Your app is running.")
	})

	// Start the server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
