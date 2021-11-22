package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/rammstein4o/15puzzle/game"
)

const (
	appID        = "com.github.rammstein4o.15puzzle"
	appVersion   = "0.0.1"
	appVersionID = "1"
	appName      = "15 Puzzle"
)

//go:generate fyne bundle -o icon.go Icon.png

func main() {
	app := app.NewWithID(appID)
	app.SetIcon(resourceIconPng)

	title := app.Preferences().StringWithFallback("windowTitle", appName)
	win := app.NewWindow(title)

	g := game.NewGame(app, win)

	win.SetContent(g)
	win.ShowAndRun()
}
