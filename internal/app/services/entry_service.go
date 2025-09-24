package services

import (
	"context"
	"diary-app/internal/domain/entities"
	"diary-app/internal/domain/repo"
)

type EntryService struct {
	repo repo.EntryRepo
}

func NewEntryService(r repo.EntryRepo) *EntryService {
	return &EntryService{repo: r}
}

func (es *EntryService) CreateEntry(ctx context.Context, entry entities.Entry) int {
	entryID := es.repo.AddEntry(ctx, entry)
	return entryID
}

func (es *EntryService) UpdateEntry(ctx context.Context, entry entities.Entry) {
	es.repo.UpdateEntry(ctx, entry)
}

func (es *EntryService) GetAllEntries(ctx context.Context) []*entities.Entry {
	entries, _ := es.repo.GetEntries(ctx)

	return entries
}

func (es *EntryService) GetEntry(ctx context.Context, id int) (*entities.Entry, error) {
	entry, err := es.repo.GetEntry(ctx, id)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (es *EntryService) DeleteEntrySrc(ctx context.Context, id int) {
	es.repo.DeleteEntry(ctx, id)
}
