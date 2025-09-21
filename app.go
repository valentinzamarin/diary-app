package main

import (
	"context"
	"diary-app/internal/app/services"
	"diary-app/internal/domain/entities"
	"time"
)

type App struct {
	ctx          context.Context
	entryService *services.EntryService
}

func NewApp(service *services.EntryService) *App {
	return &App{
		entryService: service,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CreateEntry(title, content, dateStr string) {
	date, _ := time.Parse(time.RFC3339, dateStr)

	newEntry := entities.Entry{
		Title: title,
		Text:  content,
		Date:  date,
	}

	a.entryService.CreateEntry(a.ctx, newEntry)
}
