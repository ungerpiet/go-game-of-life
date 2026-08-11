package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowTitle("Conway's Game of Life")

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
