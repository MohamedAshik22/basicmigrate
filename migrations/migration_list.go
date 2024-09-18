package migrations

// Import the migrations (without importing the db package to avoid cyclic imports)
import (
	"gorm.io/gorm"
)

// Migration interface is defined in the db package
type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
	MigrationID() string
}

// List of migrations
func GetMigrations() []Migration {
	return []Migration{
		CreateMigrationHistoriesTable{},
		CreateUserTable{},
		CreatePostTable{},
		CreateCommentTable{},
		AddAgeToUsers{},
		AddDobToUsers{},
	}
}
