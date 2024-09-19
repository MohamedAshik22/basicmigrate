package migrations

import (
	"basicmigrate/models"

	"gorm.io/gorm"
)

type CreateMigrationHistoriesTable struct{}

func (m CreateMigrationHistoriesTable) Up(db *gorm.DB) error {
	return db.Migrator().CreateTable(models.MigrationHistory{})
}

func (m CreateMigrationHistoriesTable) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(models.MigrationHistory{})
}

func (m CreateMigrationHistoriesTable) MigrationID() string {
	return "0000_create_migration_histories"
}
