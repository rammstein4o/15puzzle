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
	widget.BaseWidget
	game    *Game
	seconds time.Duration
	txt     *canvas.Text
	ticker  ticker.Interface
	objects []fyne.CanvasObject
}

func (t *timer) init() *timer {
	t.ExtendBaseWidget(t)
	t.seconds = 0
	t.txt = canvas.NewText("00:00:00", color.White)
	t.ticker = ticker.NewDefaultTicker()

	t.objects = []fyne.CanvasObject{
		container.NewCenter(
			t.txt,
		),
	}

	return t
}

func (t *timer) increment() {
	t.seconds += 1 * time.Second
	t.Refresh()
}

func (t *timer) CreateRenderer() fyne.WidgetRenderer {
	return t
}

func (t *timer) Destroy() {}

func (t *timer) Layout(s fyne.Size) {
	t.objects[0].Resize(s)
}

func (t *timer) MinSize() fyne.Size {
	return fyne.NewSize(gameMinWidth/3, toolbarHeight)
}

func (t *timer) Objects() []fyne.CanvasObject {
	return t.objects
}

func (t *timer) Refresh() {
	t.txt.Text = utils.FormatDuration(t.seconds)
	t.txt.Refresh()
}

func (t *timer) Reset() {
	t.seconds = 0
	t.Refresh()
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

func (t *timer) IsPaused() bool {
	return t.ticker.IsPaused()
}

func newTimer(game *Game) *timer {
	t := &timer{game: game}
	return t.init()
}
