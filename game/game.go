package game

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/rammstein4o/15puzzle/utils"
)

type Game struct {
	widget.BaseWidget
	objects []fyne.CanvasObject
	win     fyne.Window
	puzzle  [16]uint8
	board   *board
	toolbar *toolbar
	timer   *timer
}

func (g *Game) init() *Game {
	g.ExtendBaseWidget(g)
	g.Shuffle()
	g.timer = newTimer(g)
	g.toolbar = newToolbar(g)
	g.board = newBoard(g)
	g.objects = []fyne.CanvasObject{
		g.toolbar,
		g.board,
	}

	return g
}

func (g *Game) CreateRenderer() fyne.WidgetRenderer {
	return g
}

func (g *Game) Destroy() {}

func (g *Game) Layout(s fyne.Size) {
	g.toolbar.Resize(fyne.NewSize(s.Width, 50))
	g.toolbar.Move(fyne.NewPos(0, 0))

	g.board.Resize(fyne.NewSize(s.Width, s.Height-50))
	g.board.Move(fyne.NewPos(0, 50))
}

func (g *Game) MinSize() fyne.Size {
	return fyne.NewSize(240, 290)
}

func (g *Game) Objects() []fyne.CanvasObject {
	return g.objects
}

func (g *Game) Refresh() {
}

func (g *Game) Shuffle() {
	tmp := [16]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(tmp), func(i, j int) { tmp[i], tmp[j] = tmp[j], tmp[i] })
		if utils.IsSolvable(tmp[:]) {
			break
		}
	}

	g.puzzle = tmp
}

func (g *Game) Check() {
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

func (g *Game) SwitchItems(src, dst int) {
	g.puzzle[src], g.puzzle[dst] = g.puzzle[dst], g.puzzle[src]
}

func NewGame(win fyne.Window) *Game {
	g := &Game{
		win: win,
	}

	return g.init()
}
