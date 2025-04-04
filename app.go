package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (b *App) shutdown(ctx context.Context) {

}

func (b *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	_, _ = runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:          runtime.ErrorDialog,
		Title:         "Application Error",
		Message:       "Only one instance of the application is allowed. Aborting.",
		DefaultButton: "OK",
	})
}
