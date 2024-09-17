package main

import (
	"basicmigrate/db"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	migrateUp := flag.Bool("up", false, "Apply all migrations")
	migrateDown := flag.Bool("down", false, "Rollback all migrations")
	flag.Parse()

	// Connect to the database
	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if *migrateUp {
		fmt.Println("Running migrations...")
		db.RunMigrations(dbConn)
	} else if *migrateDown {
		fmt.Println("Rolling back migrations...")
		db.RollbackMigrations(dbConn)
	} else {
		fmt.Println("Starting the app...")

		// Define a simple HTTP server (for demonstration purposes)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World! Your app is running!")
		})

		// Start the HTTP server
		log.Println("Server is running at http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Failed to start the server: %v", err)
		}
	}

}
