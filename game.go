package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//go:generate fyne bundle -o icon.go Icon.png

type game struct {
	app    fyne.App
	win    fyne.Window
	puzzle [16]uint8
	board  *board
}

func (g *game) Shuffle() {
	tmp := [16]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tmp), func(i, j int) { tmp[i], tmp[j] = tmp[j], tmp[i] })
	g.puzzle = tmp
}

func (g *game) Check() {
	prev := uint8(0)
	for i := 0; i < 15; i++ {
		if prev < g.puzzle[i] {
			prev = g.puzzle[i]
		} else {
			return
		}
	}

	dialog.ShowInformation("Success", "You solved the puzzle!", g.win)
}

func (g *game) SwitchItems(src, dst int) {
	g.puzzle[src], g.puzzle[dst] = g.puzzle[dst], g.puzzle[src]
}

func (g *game) Run() {
	g.Shuffle()

	g.board = newBoard(g)

	g.win.SetContent(
		container.NewVBox(
			widget.NewToolbar(
				widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
					g.Shuffle()
					g.board.Refresh()
				}),
				widget.NewToolbarAction(theme.LogoutIcon(), func() {
					g.win.Close()
				}),
			),
			container.NewMax(
				g.board,
			),
		),
	)

	g.win.ShowAndRun()
}

func newGame() *game {
	app := app.NewWithID("15puzzle")
	app.SetIcon(resourceIconPng)

	return &game{
		app: app,
		win: app.NewWindow("15 Puzzle"),
	}
}
