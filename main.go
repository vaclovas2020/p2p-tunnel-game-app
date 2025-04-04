package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "P2P Tunnel Game",
		Width:     1024,
		Height:    768,
		MinWidth:  1024,
		MinHeight: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: options.NewRGB(255, 255, 255),
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: "github.com/vaclovas2020/p2p-tunnel-game-app",
			OnSecondInstanceLaunch: func(secondInstanceData options.SecondInstanceData) {
				_, _ = runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
					Type:          runtime.ErrorDialog,
					Title:         "Application Error",
					Message:       "Only one instance of the application is allowed. Aborting.",
					DefaultButton: "OK",
				})
			},
		},
		Bind: []any{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
