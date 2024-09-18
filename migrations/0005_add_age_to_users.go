package migrations

import (
	"basicmigrate/models" // Import the User model from your models package

	"gorm.io/gorm"
)

type AddAgeToUsers struct{}

// Up method to add the age column to the users table
func (m AddAgeToUsers) Up(db *gorm.DB) error {
	// Add the 'age' column to the 'users' table
	return db.Migrator().AddColumn(models.User{}, "Age")
}

// Down method to remove the age column if rolled back
func (m AddAgeToUsers) Down(db *gorm.DB) error {
	// Remove the 'age' column from the 'users' table
	return db.Migrator().DropColumn(models.User{}, "Age")
}

// MigrationID returns the unique ID of the migration
func (m AddAgeToUsers) MigrationID() string {
	return "0005_add_age_to_users"
}
