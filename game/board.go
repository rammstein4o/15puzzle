package game

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type board struct {
	widget.BaseWidget
	objects []fyne.CanvasObject
	game    *Game
	grid    *fyne.Container
}

func (b *board) init() *board {
	b.ExtendBaseWidget(b)

	bg := canvas.NewRectangle(theme.ForegroundColor())
	b.grid = container.NewGridWithColumns(b.game.puzzleCols)

	b.objects = []fyne.CanvasObject{
		container.NewMax(bg, container.NewPadded(b.grid)),
	}

	return b
}

func (b *board) CreateRenderer() fyne.WidgetRenderer {
	return b
}

func (b *board) Destroy() {}

func (b *board) Layout(s fyne.Size) {
	b.objects[0].Resize(s)
	b.Refresh()
}

func (b *board) MinSize() fyne.Size {
	return fyne.NewSize(gameMinWidth, gameMinWidth)
}

func (b *board) Objects() []fyne.CanvasObject {
	return b.objects
}

func (b *board) Refresh() {
	b.grid.Objects = []fyne.CanvasObject{}
	for i := 0; i < b.game.puzzleLen; i++ {
		b.grid.Add(newTile(b, b.game.seed[i]))
	}
	b.grid.Refresh()
}

func (b *board) MoveTile(t *tile) {
	if t.val == 0 {
		return
	}
	if b.game.timer.IsPaused() {
		return
	}

	var src, dst int
	for idx, val := range b.game.seed {
		if val == t.val {
			src = idx
		}
		if val == 0 {
			dst = idx
		}
	}

	if b.isMovePossible(src, dst) {
		b.switchItems(src, dst)
		b.game.SwitchItems(src, dst)
	}

	b.game.Check()
}

func (b *board) isMovePossible(src, dst int) bool {
	srcCol := src % b.game.puzzleCols
	srcRow := src / b.game.puzzleCols
	dstCol := dst % b.game.puzzleCols
	dstRow := dst / b.game.puzzleCols

	if math.Abs(float64(srcCol)-float64(dstCol)) > 1 || math.Abs(float64(srcRow)-float64(dstRow)) > 1 {
		return false
	}
	if srcCol == dstCol || srcRow == dstRow {
		return true
	}
	return false
}

func (b *board) switchItems(src, dst int) {
	b.grid.Objects[src], b.grid.Objects[dst] = b.grid.Objects[dst], b.grid.Objects[src]
	b.grid.Refresh()
}

func newBoard(g *Game) *board {
	b := &board{
		game: g,
	}

	return b.init()
}
