package game

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type tile struct {
	widget.BaseWidget
	board   *board
	val     uint8
	objects []fyne.CanvasObject
}

func (t *tile) CreateRenderer() fyne.WidgetRenderer {
	return t
}

func (t *tile) Destroy() {}

func (t *tile) Layout(s fyne.Size) {
	t.objects[0].Resize(s)
}

func (t *tile) MinSize() fyne.Size {
	size := gameMinWidth / float32(t.board.game.puzzleCols)
	return fyne.NewSize(size, size)
}

func (t *tile) Objects() []fyne.CanvasObject {
	return t.objects
}

func (t *tile) Refresh() {}

func (t *tile) Tapped(_ *fyne.PointEvent) {
	t.board.MoveTile(t)
}

func newTile(brd *board, val uint8) *tile {
	var objects []fyne.CanvasObject

	if val != 0 {
		txt := canvas.NewText(fmt.Sprintf("%v", val), theme.ForegroundColor())
		txt.TextStyle = fyne.TextStyle{
			Bold: true,
		}

		objects = append(
			objects,
			canvas.NewRectangle(brd.game.theme.TileColor()),
			container.NewCenter(txt),
		)
	}

	ret := &tile{
		board: brd,
		val:   val,
		objects: []fyne.CanvasObject{
			container.NewMax(objects...),
		},
	}
	ret.ExtendBaseWidget(ret)
	return ret
}
