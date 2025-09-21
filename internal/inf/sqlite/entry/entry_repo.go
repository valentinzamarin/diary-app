package entry

import (
	"context"
	"diary-app/internal/domain/entities"

	"gorm.io/gorm"
)

type GormEntriesRepo struct {
	db *gorm.DB
}

func NewGormEntriesRepo(db *gorm.DB) *GormEntriesRepo {
	return &GormEntriesRepo{
		db: db,
	}
}

func (ep *GormEntriesRepo) AddEntry(ctx context.Context, entry entities.Entry) {
	e := ToEntryModel(entry)
	ep.db.WithContext(ctx).Create(e)
}

func (ep *GormEntriesRepo) GetEntries(ctx context.Context) ([]*entities.Entry, error) {
	var dbEntryModels []EntryModel
	if err := ep.db.Find(&dbEntryModels).Error; err != nil {
		return nil, err
	}
	entries := make([]*entities.Entry, len(dbEntryModels))
	for i, entry := range dbEntryModels {

		entries[i] = ToEntry(&entry)
	}
	return entries, nil

}
