package entry

import (
	"context"
	"diary-app/internal/domain/entities"
	"errors"
	"fmt"

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

func (ep *GormEntriesRepo) GetEntry(ctx context.Context, id int) (*entities.Entry, error) {
	var em EntryModel

	result := ep.db.WithContext(ctx).
		Where("id = ?", id).
		First(&em)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("запись с ID %d не найдена", id)
		}
		return nil, fmt.Errorf("ошибка при получении записи: %v", result.Error)
	}

	entry := ToEntry(&em)

	return entry, nil
}
