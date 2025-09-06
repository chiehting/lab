package main

import (
	"context"
	"os"
	"path/filepath"

	"howett.net/plist"
)

// App struct
type App struct {
	ctx context.Context
	url string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		url: getDefaultHomePageFromPlist(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetURL 返回URL
func (a *App) GetURL() string {
	return a.url
}

// InfoPlist read Info.plist structure
type InfoPlist struct {
	DefaultHomePage string `plist:"DefaultHomePage"`
}

// 從 Info.plist 讀取預設首頁
func getDefaultHomePageFromPlist() string {
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}

	// 如果是 .app bundle，Info.plist 在 ../Info.plist
	plistPath := filepath.Join(filepath.Dir(execPath), "..", "Info.plist")

	// 讀取 plist 檔案
	file, err := os.Open(plistPath)
	if err != nil {
		return ""
	}
	defer file.Close()

	var info InfoPlist
	decoder := plist.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		return ""
	}

	return info.DefaultHomePage
}
