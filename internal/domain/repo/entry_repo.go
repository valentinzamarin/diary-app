package repo

import (
	"context"
	"diary-app/internal/domain/entities"
)

type EntryRepo interface {
	AddEntry(ctx context.Context, entry entities.Entry)
	GetEntries(ctx context.Context) ([]*entities.Entry, error)
	GetEntry(ctx context.Context, id int) (*entities.Entry, error)
}
