package migrations

import (
	"gorm.io/gorm"
)

type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
	MigrationID() string
}

func GetMigrations() []Migration {
	return []Migration{
		CreateMigrationHistoriesTable{},
		CreateUserTable{},
		CreatePostTable{},
		CreateCommentTable{},
		AddAgeToUsers{},
		AddDobToUsers{},
		RenameUserNameToFullName{},
	}
}
