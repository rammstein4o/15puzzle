package game

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var possibleMoves map[int][]int = map[int][]int{
	0:  {1, 4},
	1:  {0, 2, 5},
	2:  {1, 3, 6},
	3:  {2, 7},
	4:  {0, 5, 8},
	5:  {1, 4, 6, 9},
	6:  {2, 5, 7, 10},
	7:  {3, 6, 11},
	8:  {4, 9, 12},
	9:  {5, 8, 10, 13},
	10: {6, 9, 11, 14},
	11: {7, 10, 15},
	12: {8, 13},
	13: {9, 12, 14},
	14: {10, 13, 15},
	15: {11, 14},
}

type board struct {
	widget.BaseWidget
	objects []fyne.CanvasObject
	game    *Game
	grid    *fyne.Container
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
	return fyne.NewSize(240, 240)
}

func (b *board) Objects() []fyne.CanvasObject {
	return b.objects
}

func (b *board) Refresh() {
	b.grid.Objects = []fyne.CanvasObject{}
	for i := 0; i <= 15; i++ {
		b.grid.Add(newTile(b, b.game.puzzle[i]))
	}
	b.grid.Refresh()
}

func (b *board) MoveTile(t *tile) {
	if t.val == 0 {
		return
	}

	var src, dst int
	for idx, val := range b.game.puzzle {
		if val == t.val {
			src = idx
		}
		if val == 0 {
			dst = idx
		}
	}

	if b.isMovePossible(src, dst) {
		// Switch tiles in grid
		b.grid.Objects[src], b.grid.Objects[dst] = b.grid.Objects[dst], b.grid.Objects[src]
		b.grid.Refresh()

		// Switch tiles in puzzle
		b.game.SwitchItems(src, dst)
	}

	b.game.Check()
}

func (b *board) isMovePossible(src, dst int) bool {
	if moves, found := possibleMoves[src]; found {
		for _, move := range moves {
			if move == dst {
				return true
			}
		}
	}
	return false
}

func newBoard(g *Game) *board {
	bg := canvas.NewRectangle(color.White)
	grid := container.NewGridWithColumns(4)

	// for i := 0; i <= 15; i++ {
	// 	grid.Add(newTile(g, g.puzzle[i]))
	// }

	c := &board{
		objects: []fyne.CanvasObject{
			container.NewMax(bg, container.NewPadded(grid)),
		},
		game: g,
		grid: grid,
	}

	c.ExtendBaseWidget(c)
	return c
}
