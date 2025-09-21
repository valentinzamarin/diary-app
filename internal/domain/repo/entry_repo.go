package repo

import (
	"context"
	"diary-app/internal/domain/entities"
)

type EntryRepo interface {
	AddEntry(ctx context.Context, entry entities.Entry)
}
