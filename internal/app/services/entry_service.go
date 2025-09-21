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
