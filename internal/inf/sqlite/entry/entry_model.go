package entry

import "time"

type EntryModel struct {
	ID    uint      `gorm:"primaryKey;autoIncrement"`
	Title string    `gorm:"not null;"`
	Text  string    `gorm:"not null"`
	Date  time.Time `gorm:"not null;index"`
}

func (EntryModel) TableName() string {
	return "entries"
}
