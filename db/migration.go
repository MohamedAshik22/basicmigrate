package db

import (
	"basicmigrate/migrations"
	"log"

	"gorm.io/gorm"
)

// Migration interface defining Up and Down methods
type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

// RunMigrations executes all migrations' Up functions
func RunMigrations(db *gorm.DB) {
	migrationList := migrations.GetMigrations()
	for _, migration := range migrationList {
		if err := migration.Up(db); err != nil {
			log.Fatalf("Failed to apply migration: %v", err)
		}
	}
	log.Println("All migrations applied successfully!")
}

// RollbackMigrations executes all migrations' Down functions in reverse order
func RollbackMigrations(db *gorm.DB) {
	migrationList := migrations.GetMigrations()
	for i := len(migrationList) - 1; i >= 0; i-- {
		if err := migrationList[i].Down(db); err != nil {
			log.Fatalf("Failed to rollback migration: %v", err)
		}
	}
	log.Println("All migrations rolled back successfully!")
}
