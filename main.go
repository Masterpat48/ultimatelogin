package main

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed db.json
var embeddedDB []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "ultimatelogin",
		Width:     1000,
		MinWidth:  700,
		Height:    700,
		MinHeight: 400,
		Frameless: true,
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

func (a *App) initDatabase() error {
	// Percorso per il file database nella home dell'utente
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Creo una cartella nascosta per l'app
	appDir := filepath.Join(homeDir, ".ultimatelogin")
	dbPath := filepath.Join(appDir, "db.json")

	// Crea la directory se non esiste
	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		return err
	}

	// Se il file non esiste, crealo con i dati embedded
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err = os.WriteFile(dbPath, embeddedDB, 0644)
		if err != nil {
			return err
		}
	}

	// Ora usa dbPath per le operazioni sul database
	a.dbPath = dbPath
	return nil
}
