package utils

import (
	"image/color"

	"fyne.io/fyne/v2"
	fyneTheme "fyne.io/fyne/v2/theme"
)

const (
	ColorNameTile fyne.ThemeColorName = "tile"
)

type ThemeInteface interface {
	fyne.Theme
	TileColor() color.Color
}

type theme struct{}

func (t *theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == ColorNameTile {
		if variant == fyneTheme.VariantLight {
			return color.RGBA{0xcc, 0xcc, 0xcc, 0xFF}
		}
		return color.RGBA{0x22, 0x66, 0x66, 0xFF}
	}
	return fyneTheme.DefaultTheme().Color(name, variant)
}

func (t *theme) Font(style fyne.TextStyle) fyne.Resource {
	return fyneTheme.DefaultTheme().Font(style)
}

func (t *theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return fyneTheme.DefaultTheme().Icon(name)
}

func (t *theme) Size(name fyne.ThemeSizeName) float32 {
	return fyneTheme.DefaultTheme().Size(name)
}

func (t *theme) TileColor() color.Color {
	return t.Color(ColorNameTile, fyne.CurrentApp().Settings().ThemeVariant())
}

func NewTheme() ThemeInteface {
	return &theme{}
}
