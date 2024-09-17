package migrations

import "gorm.io/gorm"

type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

// GetMigrations returns a list of all the migrations to be applied
func GetMigrations() []Migration {
	return []Migration{
		CreateUserTable{},
		CreatePostTable{},
		CreateCommentTable{},
	}
}
