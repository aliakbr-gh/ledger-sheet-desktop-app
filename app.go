package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	goruntime "runtime"

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

	runtime.WindowMaximise(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SavePDF(fileName string, data []byte) error {

	var basePath string

	switch goruntime.GOOS {

	case "windows":
		customPath := `D:\Khawaj Muhammad Khan\Shop Documents\Daily Worksheet 2026`

		// Check if custom path exists
		if _, err := os.Stat(customPath); err == nil {
			basePath = customPath
		} else {
			// Fallback to Downloads folder
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			basePath = filepath.Join(home, "Downloads", "KMKCommunication")
		}

	case "darwin", "linux":
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		basePath = filepath.Join(home, "Downloads", "KMKCommunication")

	default:
		home, _ := os.UserHomeDir()
		basePath = filepath.Join(home, "KMKCommunication")
	}

	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(basePath, fileName)

	return os.WriteFile(fullPath, data, 0644)
}