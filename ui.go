package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type button struct {
	x, y, width, height int
	label               string
	backgroundColor     color.Color
}

func (b *button) IsClicked(mouseX, mouseY int) bool {
	if mouseX >= b.x && mouseX <= b.x+b.width && mouseY >= b.y && mouseY <= b.y+b.height {
		return true
	}
	return false
}

func drawButton(b button, screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	btnColor := b.backgroundColor

	if b.IsClicked(mx, my) {
		btnColor = color.RGBA{65, 71, 83, 255}
	}

	vector.FillRect(screen, float32(b.x), float32(b.y), float32(b.width), float32(b.height), btnColor, true)
	ebitenutil.DebugPrintAt(screen, b.label, b.x+4, b.y+2)
}
