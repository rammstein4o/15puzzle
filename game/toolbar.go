package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type toolbar struct {
	widget.BaseWidget
	game       *Game
	objects    []fyne.CanvasObject
	refreshBtn *widget.Button
	playBtn    *widget.Button
	pauseBtn   *widget.Button
	prefsBtn   *widget.Button
}

func (t *toolbar) init() *toolbar {
	t.ExtendBaseWidget(t)

	t.prefsBtn = widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		var newPuzzleType PuzzleType

		gameTypeSelect := widget.NewFormItem(
			"Game Type",
			widget.NewSelect(
				[]string{"8 Puzzle", "15 Puzzle", "24 Puzzle", "35 Puzzle"},
				func(selected string) {
					switch selected {
					case "8 Puzzle":
						newPuzzleType = Puzzle8
					case "24 Puzzle":
						newPuzzleType = Puzzle24
					case "35 Puzzle":
						newPuzzleType = Puzzle35
					case "15 Puzzle":
						fallthrough
					default:
						newPuzzleType = Puzzle15
					}
				},
			),
		)

		dialog.ShowForm(
			"Preferences",
			"Save",
			"Cancel",
			[]*widget.FormItem{gameTypeSelect},
			func(res bool) {
				if res {
					t.game.SwitchPuzzleType(newPuzzleType)
				}
			},
			t.game.win,
		)
	})
	t.prefsBtn.Importance = widget.LowImportance

	t.refreshBtn = widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		t.game.Reset()
	})
	t.refreshBtn.Importance = widget.LowImportance

	t.playBtn = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		t.game.timer.Resume()
		t.Refresh()
	})
	t.playBtn.Importance = widget.LowImportance

	t.pauseBtn = widget.NewButtonWithIcon("", theme.MediaPauseIcon(), func() {
		t.game.timer.Pause()
		t.Refresh()
	})
	t.pauseBtn.Importance = widget.LowImportance

	t.objects = []fyne.CanvasObject{
		container.NewHBox(
			t.prefsBtn,
			t.refreshBtn,
			t.playBtn,
			t.pauseBtn,
			t.game.timer,
		),
	}

	t.Refresh()

	return t
}

func (t *toolbar) CreateRenderer() fyne.WidgetRenderer {
	return t
}

func (t *toolbar) Destroy() {}

func (t *toolbar) Layout(s fyne.Size) {
	t.objects[0].Resize(s)
}

func (t *toolbar) MinSize() fyne.Size {
	return fyne.NewSize(gameMinWidth, toolbarHeight)
}

func (t *toolbar) Objects() []fyne.CanvasObject {
	return t.objects
}

func (t *toolbar) Refresh() {
	if t.game.timer.IsPaused() {
		t.playBtn.Show()
		t.pauseBtn.Hide()
	} else {
		t.playBtn.Hide()
		t.pauseBtn.Show()
	}
	if t.game.solved {
		t.playBtn.Hide()
	}
}

func newToolbar(game *Game) *toolbar {
	t := &toolbar{game: game}
	return t.init()
}
