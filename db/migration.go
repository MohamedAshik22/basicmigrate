package db

import (
	"basicmigrate/migrations"
	"basicmigrate/models"
	"log"
	"time"

	"gorm.io/gorm"
)

type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
	MigrationID() string
}

func RunMigrations(db *gorm.DB) {
	migrationList := migrations.GetMigrations()

	for _, migration := range migrationList {
		if !isMigrationApplied(db, migration.MigrationID()) {
			log.Printf("Applying migration: %s", migration.MigrationID())
			if err := migration.Up(db); err != nil {
				log.Fatalf("Failed to apply migration: %v", err)
			}
			recordMigration(db, migration.MigrationID())
		} else {
			log.Printf("Migration %s has already been applied", migration.MigrationID())
		}
	}
	log.Println("Migrations ran successfully!")
}

func RollbackMigrations(db *gorm.DB) {
	migrationList := migrations.GetMigrations()
	for i := len(migrationList) - 1; i >= 0; i-- {
		migration := migrationList[i]
		if isMigrationApplied(db, migration.MigrationID()) {
			log.Printf("Rolling back migration: %s", migration.MigrationID())
			if err := migration.Down(db); err != nil {
				log.Fatalf("Failed to rollback migration: %v", err)
			}
			removeMigrationRecord(db, migration.MigrationID())
		}
	}
	log.Println("Migrations rolled back successfully!")
}

func RollbackMigration(db *gorm.DB, migrationID string) {
	migrationList := migrations.GetMigrations()
	for _, migration := range migrationList {
		if migration.MigrationID() == migrationID {
			if isMigrationApplied(db, migrationID) {
				log.Printf("Rolling back migration: %s", migrationID)
				if err := migration.Down(db); err != nil {
					log.Fatalf("Failed to rollback migration: %v", err)
				}
				removeMigrationRecord(db, migrationID)
				log.Printf("Migration %s rolled back successfully", migrationID)
			} else {
				log.Printf("Migration %s has not been applied, skipping rollback", migrationID)
			}
			return
		}
	}
	log.Printf("Migration %s not found", migrationID)
}

func isMigrationApplied(db *gorm.DB, migrationID string) bool {
	var count int64
	db.Model(&models.MigrationHistory{}).Where("migration_id = ?", migrationID).Count(&count)
	return count > 0
}

func recordMigration(db *gorm.DB, migrationID string) {
	db.Create(&models.MigrationHistory{
		MigrationID: migrationID,
		AppliedAt:   time.Now(),
	})
}

func removeMigrationRecord(db *gorm.DB, migrationID string) {
	db.Where("migration_id = ?", migrationID).Delete(&models.MigrationHistory{})
}
