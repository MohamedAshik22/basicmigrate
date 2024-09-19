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
	specificDown := flag.String("undo", "", "Undo specific migration by ID (e.g., '0001_create_users')")
	flag.Parse()

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if *migrateUp {
		fmt.Println("Running migrations...")
		db.RunMigrations(dbConn)
		return
	} else if *migrateDown {
		fmt.Println("Rolling back all migrations...")
		db.RollbackMigrations(dbConn)
		return
	} else if *specificDown != "" {
		fmt.Printf("Rolling back migration: %s\n", *specificDown)
		db.RollbackMigration(dbConn, *specificDown)
		return
	}

	fmt.Println("Starting server on port 8080...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Your app is running.")
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
