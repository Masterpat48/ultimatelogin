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
	// databse dir connection
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Creating the hidden folder
	appDir := filepath.Join(homeDir, ".ultimatelogin")
	dbPath := filepath.Join(appDir, "db.json")

	// If the dir doesn't exist we create it here
	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		return err
	}

	// If the file doesnt't exist we create it with embedded data
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err = os.WriteFile(dbPath, embeddedDB, 0644)
		if err != nil {
			return err
		}
	}

	// Using the dit for the database
	a.dbPath = dbPath
	return nil
}
