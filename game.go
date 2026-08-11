package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const screenWidth = 320
const screenHeight = 240
const cellSize = 4

type Game struct {
	backgroundColor color.RGBA
	grid            [][]bool
	paused          bool
	pauseButton     button
	resetButton     button
	randomButton    button
}

func newGame() *Game {
	cols := screenWidth / cellSize
	rows := screenHeight / cellSize

	grid := make([][]bool, cols)

	for x := 0; x < cols; x++ {
		grid[x] = make([]bool, rows)
		for y := 0; y < rows; y++ {
			grid[x][y] = rand.Float64() < 0.2
		}
	}

	return &Game{
		backgroundColor: color.RGBA{24, 24, 30, 255},
		grid:            grid,
		paused:          true,
		pauseButton:     button{5, 5, 50, 20, "Pause", color.RGBA{40, 44, 52, 255}},
		resetButton:     button{60, 5, 50, 20, "Reset", color.RGBA{40, 44, 52, 255}},
		randomButton:    button{115, 5, 50, 20, "Random", color.RGBA{40, 44, 52, 255}},
	}
}

func (g *Game) Update() error {
	cols := len(g.grid)
	rows := len(g.grid[0])
	mouseX, mouseY := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.pauseButton.IsClicked(mouseX, mouseY) {
			g.paused = !g.paused
			if g.paused {
				g.pauseButton.label = "Play"
			} else {
				g.pauseButton.label = "Pause"
			}
			return nil
		}
		if g.resetButton.IsClicked(mouseX, mouseY) {
			for x := 0; x < cols; x++ {
				for y := 0; y < rows; y++ {
					g.grid[x][y] = false
				}
			}
			return nil
		}
		if g.randomButton.IsClicked(mouseX, mouseY) {
			for x := 0; x < cols; x++ {
				for y := 0; y < rows; y++ {
					g.grid[x][y] = rand.Float64() < 0.2
				}
			}
			return nil
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = !g.paused
	}

	if !g.paused {
		nextGrid := make([][]bool, cols)
		for x := 0; x < cols; x++ {
			nextGrid[x] = make([]bool, rows)
			for y := 0; y < rows; y++ {
				nextGrid[x][y] = g.nextCellState(x, y, g.grid[x][y])
			}
		}
		g.grid = nextGrid
	} else {
		isOnAnyButton := g.pauseButton.IsClicked(mouseX, mouseY) ||
			g.resetButton.IsClicked(mouseX, mouseY) ||
			g.randomButton.IsClicked(mouseX, mouseY)

		if !isOnAnyButton {
			gridX := mouseX / cellSize
			gridY := mouseY / cellSize
			if gridX >= 0 && gridX < cols && gridY >= 0 && gridY < rows {
				if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
					g.grid[gridX][gridY] = true
				}
				if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
					g.grid[gridX][gridY] = false
				}
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.backgroundColor)

	for x := 0; x < len(g.grid); x++ {
		for y := 0; y < len(g.grid[x]); y++ {
			if g.grid[x][y] {
				vector.FillRect(screen, float32(x*cellSize), float32(y*cellSize), cellSize, cellSize, color.RGBA{0, 240, 255, 255}, true)
			}
		}
	}

	drawButton(g.pauseButton, screen)
	drawButton(g.resetButton, screen)
	drawButton(g.randomButton, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) nextCellState(x int, y int, alive bool) bool {
	cols := len(g.grid)
	rows := len(g.grid[0])
	aliveNeighbors := 0

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx := (x + dx + cols) % cols
			ny := (y + dy + rows) % rows

			if g.grid[nx][ny] {
				aliveNeighbors++
			}
		}
	}

	if alive {
		return aliveNeighbors == 2 || aliveNeighbors == 3
	}
	return aliveNeighbors == 3
}
