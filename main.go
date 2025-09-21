package main

import (
	"diary-app/internal/app/services"
	"diary-app/internal/inf/sqlite"
	"diary-app/internal/inf/sqlite/entry"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	basePath := "./notebook.db"

	db, _ := sqlite.NewConnection(basePath)
	sqlite.Migrate(db)

	entryRepo := entry.NewGormEntriesRepo(db)

	EntryService := services.NewEntryService(entryRepo)

	app := NewApp(EntryService)

	err := wails.Run(&options.App{
		Title:  "diary-app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
