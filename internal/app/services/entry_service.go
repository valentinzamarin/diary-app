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

func (es *EntryService) CreateEntry(ctx context.Context, entry entities.Entry) {
	es.repo.AddEntry(ctx, entry)
}

func (es *EntryService) GetAllEntries(ctx context.Context) []*entities.Entry {
	entries, _ := es.repo.GetEntries(ctx)

	return entries
}
