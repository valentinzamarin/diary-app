package main

import (
	"context"
	"diary-app/internal/app/services"
	"diary-app/internal/domain/entities"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func (a *App) CreateEntry(title, content string) {
	created_date := time.Now()

	newEntry := entities.Entry{
		Title: title,
		Text:  content,
		Date:  created_date,
	}
	// need
	// cuz instead
	// get ID 0
	// @ front
	id := a.entryService.CreateEntry(a.ctx, newEntry)
	newEntry.ID = id

	runtime.EventsEmit(a.ctx, "entry:created", newEntry)
}

func (a *App) UpdateEntry(updEntrу entities.Entry) {
	a.entryService.UpdateEntry(a.ctx, updEntrу)
}

func (a *App) GetEntries() []*entities.Entry {
	return a.entryService.GetAllEntries(a.ctx)
}

func (a *App) GetEntry(id int) (*entities.Entry, error) {
	content, err := a.entryService.GetEntry(a.ctx, id)
	return content, err
}

func (a *App) DeleteEntryApp(id int) {
	a.entryService.DeleteEntrySrc(a.ctx, id)
	runtime.EventsEmit(a.ctx, "entry:deleted", id)
}
