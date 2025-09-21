package sqlite

import (
	"diary-app/internal/inf/sqlite/entry"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entry.EntryModel{},
	)
}
