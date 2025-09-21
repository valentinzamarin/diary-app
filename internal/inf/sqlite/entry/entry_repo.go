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
