package main

import (
	"dayz-tool/service"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:         "DayzTool",
		Width:         1057,
		MinWidth:      1057,
		Height:        768,
		MinHeight:     768,
		DisableResize: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Bind:             getBind(app),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func getBind(app *App) []interface{} {
	binds := []interface{}{}
	binds = append(binds, app)
	binds = append(binds, service.InitService()...)
	return binds
}
