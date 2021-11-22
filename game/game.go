package game

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/rammstein4o/15puzzle/utils"
)

type PuzzleType uint8

const (
	Puzzle8 PuzzleType = iota
	Puzzle15
	Puzzle24
	Puzzle35
)

const (
	gameMinWidth  float32 = 300
	toolbarHeight float32 = 40
)

type Game struct {
	widget.BaseWidget
	objects    []fyne.CanvasObject
	app        fyne.App
	win        fyne.Window
	puzzleType PuzzleType
	puzzleLen  int
	puzzleCols int
	seed       []uint8
	board      *board
	toolbar    *toolbar
	timer      *timer
	solved     bool
}

func (g *Game) init() *Game {
	g.ExtendBaseWidget(g)

	pt := PuzzleType(g.app.Preferences().IntWithFallback("gameType", int(Puzzle15)))
	g.setGameBoard(pt)
	g.setToolbar()

	g.objects = []fyne.CanvasObject{
		g.toolbar,
		g.board,
	}

	return g
}

func (g *Game) setToolbar() {
	g.timer = newTimer(g)
	g.timer.Start()
	g.toolbar = newToolbar(g)
}

func (g *Game) setGameBoard(pt PuzzleType) {
	g.puzzleType = pt
	g.Shuffle()
	g.puzzleLen = len(g.seed)
	g.puzzleCols = int(math.Sqrt(float64(g.puzzleLen)))
	g.board = newBoard(g)
}

func (g *Game) CreateRenderer() fyne.WidgetRenderer {
	return g
}

func (g *Game) Destroy() {}

func (g *Game) Layout(s fyne.Size) {
	g.toolbar.Resize(fyne.NewSize(s.Width, toolbarHeight))
	g.toolbar.Move(fyne.NewPos(0, 0))

	g.board.Resize(fyne.NewSize(s.Width, s.Height-toolbarHeight))
	g.board.Move(fyne.NewPos(0, toolbarHeight))
}

func (g *Game) MinSize() fyne.Size {
	return fyne.NewSize(gameMinWidth, gameMinWidth+toolbarHeight)
}

func (g *Game) Objects() []fyne.CanvasObject {
	return g.objects
}

func (g *Game) Refresh() {
	g.toolbar.Refresh()
	g.board.Refresh()
}

func (g *Game) Shuffle() {
	var tmp []uint8
	switch g.puzzleType {
	case Puzzle8:
		tmp = utils.GenerateSeed(9)
	case Puzzle24:
		tmp = utils.GenerateSeed(25)
	case Puzzle35:
		tmp = utils.GenerateSeed(36)
	case Puzzle15:
		fallthrough
	default:
		tmp = utils.GenerateSeed(16)
	}

	for {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(tmp), func(i, j int) { tmp[i], tmp[j] = tmp[j], tmp[i] })
		if utils.IsSolvable(tmp) {
			break
		}
	}
	g.seed = tmp
	g.solved = false
}

func (g *Game) Check() {
	prev := uint8(0)
	for idx, val := range g.seed {
		if prev < val {
			prev = val
		} else {
			if val != 0 {
				return
			}
			if val == 0 && idx != g.puzzleLen-1 {
				return
			}
		}
	}

	dialog.ShowInformation("Success", "You solved the puzzle!", g.win)
	g.timer.Pause()
	g.solved = true
	g.toolbar.Refresh()
}

func (g *Game) Reset() {
	g.Shuffle()
	g.timer.Reset()
	g.timer.Resume()
	g.Refresh()
}

func (g *Game) SwitchItems(src, dst int) {
	g.seed[src], g.seed[dst] = g.seed[dst], g.seed[src]
}

func (g *Game) SwitchPuzzleType(pt PuzzleType) {
	if pt != g.puzzleType {
		// Set user preferences
		g.app.Preferences().SetInt("gameType", int(pt))
		// Set game board
		g.setGameBoard(pt)
		g.objects[1] = g.board

		// Set window title
		title := fmt.Sprintf("%d Puzzle", g.puzzleLen-1)
		g.app.Preferences().SetString("windowTitle", title)
		g.win.SetTitle(title)

		// Reset the timer
		g.timer.Reset()
	}
	g.Refresh()
}

func NewGame(app fyne.App, win fyne.Window) *Game {
	g := &Game{
		app: app,
		win: win,
	}

	return g.init()
}
