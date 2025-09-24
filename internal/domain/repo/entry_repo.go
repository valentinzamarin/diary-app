package repo

import (
	"context"
	"diary-app/internal/domain/entities"
)

/*
GetEntries - for the list
GetEntry - for the single view
*/
type EntryRepo interface {
	AddEntry(ctx context.Context, entry entities.Entry) int
	UpdateEntry(ctx context.Context, entry entities.Entry) error
	GetEntries(ctx context.Context) ([]*entities.Entry, error)
	GetEntry(ctx context.Context, id int) (*entities.Entry, error)
	DeleteEntry(ctx context.Context, id int)
}
