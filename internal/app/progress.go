package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func DrawProgressBar(screen *ebiten.Image, x, y, width, height int, value, maxValue float64, label string) {
	// Draw label above the bar
	text.Draw(screen, label, basicfont.Face7x13, x, y-3, color.Black)

	// Black outer border
	border := ebiten.NewImage(width+2, height+2)
	border.Fill(color.Black)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x-1), float64(y-1))
	screen.DrawImage(border, op)

	// Grey background
	bg := ebiten.NewImage(width, height)
	bg.Fill(color.RGBA{192, 192, 192, 255})
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(bg, op2)

	// Black fill with gap/padding
	gap := 2
	fillWidth := int(float64(width-gap*2) * (value / maxValue))
	if fillWidth > 0 {
		fill := ebiten.NewImage(fillWidth, height-gap*2)
		fill.Fill(color.Black)
		op3 := &ebiten.DrawImageOptions{}
		op3.GeoM.Translate(float64(x+gap), float64(y+gap))
		screen.DrawImage(fill, op3)
	}
}
