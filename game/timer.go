package game

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rammstein4o/15puzzle/utils"
	"github.com/rammstein4o/ticker"
)

type timer struct {
	widget.ToolbarItem
	game    *Game
	seconds time.Duration
	txt     *canvas.Text
	object  fyne.CanvasObject
	ticker  ticker.Interface
}

func (t *timer) init() *timer {
	t.seconds = 0
	t.txt = canvas.NewText("00:00:00", color.White)
	t.object = container.NewCenter(t.txt)
	t.ticker = ticker.NewDefaultTicker()
	return t
}

func (t *timer) increment() {
	t.seconds += 1 * time.Second
	t.Refresh()
}

func (t *timer) ToolbarObject() fyne.CanvasObject {
	return t.object
}

func (t *timer) Reset() {
	t.seconds = 0
	t.Refresh()
}

func (t *timer) Refresh() {
	t.txt.Text = utils.FormatDuration(t.seconds)
	t.txt.Refresh()
}

func (t *timer) Start() {
	ticks := t.ticker.Start()
	go func() {
		for range ticks {
			t.increment()
		}
	}()
}

func (t *timer) Pause() {
	t.ticker.Pause()
}

func (t *timer) Resume() {
	t.ticker.Resume()
}

func (t *timer) Stop() {
	t.ticker.Stop()
}

func newTimer(game *Game) *timer {
	t := &timer{game: game}
	return t.init()
}
