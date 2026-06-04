package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	goruntime "runtime"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowMaximise(ctx)
}

func (a *App) SavePDF(fileName string, pdfData []byte, jsonData string) error {

	var basePath string

	switch goruntime.GOOS {
	case "windows":
		customPath := `D:\Khawaj Muhammad Khan\Shop Documents\Daily Worksheet 2026`
		if _, err := os.Stat(customPath); err == nil {
			basePath = customPath
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			basePath = filepath.Join(home, "Downloads", "KMKCommunication")
		}

	default:
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		basePath = filepath.Join(home, "Downloads", "KMKCommunication")
	}

	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return err
	}

	pdfFullPath := filepath.Join(basePath, fileName)
	finalPdfPath, err := a.saveWithDuplicateHandling(pdfFullPath, pdfData)
	if err != nil {
		return err
	}

	jsonFileName := fileName[:len(fileName)-4] + ".json"
	jsonFullPath := filepath.Join(basePath, jsonFileName)
	_, err = a.saveWithDuplicateHandling(jsonFullPath, []byte(jsonData))
	if err != nil {
		return err
	}

	fmt.Printf("✅ Saved: %s\n   and   %s\n", filepath.Base(finalPdfPath), jsonFileName)
	return nil
}

func (a *App) saveWithDuplicateHandling(fullPath string, data []byte) (string, error) {

	if _, err := os.Stat(fullPath); err == nil {
		ext := filepath.Ext(fullPath)
		nameWithoutExt := fullPath[:len(fullPath)-len(ext)]

		counter := 1
		for {
			newPath := fmt.Sprintf("%s (%d)%s", nameWithoutExt, counter, ext)

			if _, err := os.Stat(newPath); os.IsNotExist(err) {
				fullPath = newPath
				break
			}
			counter++
		}
	}

	err := os.WriteFile(fullPath, data, 0644)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}