package game

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type toolbar struct {
	widget.Toolbar
	game       *Game
	refreshBtn widget.ToolbarItem
	playBtn    widget.ToolbarItem
	pauseBtn   widget.ToolbarItem
	logoutBtn  widget.ToolbarItem
}

func (t *toolbar) init() *toolbar {
	t.ExtendBaseWidget(t)

	t.logoutBtn = widget.NewToolbarAction(theme.LogoutIcon(), func() {
		t.game.win.Close()
	})

	t.Append(t.logoutBtn)

	t.refreshBtn = widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
		t.game.Shuffle()
		t.game.board.Refresh()
	})

	t.Append(t.refreshBtn)

	t.playBtn = widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
		t.game.timer.Start()
	})

	t.Append(t.playBtn)

	t.pauseBtn = widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
		t.game.timer.Pause()
	})

	t.Append(t.pauseBtn)

	t.Append(
		t.game.timer,
	)

	return t
}

func newToolbar(game *Game) *toolbar {
	t := &toolbar{game: game}
	return t.init()
}
