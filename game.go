package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const screenWidth = 320
const screenHeight = 240
const cellSize = 4

type Game struct {
	backgroundColor color.RGBA
	grid            [][]bool
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
		backgroundColor: color.RGBA{0, 0, 255, 255},
		grid:            grid,
	}
}

func (g *Game) Update() error {
	cols := len(g.grid)
	rows := len(g.grid[0])

	nextGrid := make([][]bool, cols)
	for x := 0; x < cols; x++ {
		nextGrid[x] = make([]bool, rows)
		for y := 0; y < rows; y++ {
			nextGrid[x][y] = g.nextCellState(x, y, g.grid[x][y])
		}
	}
	g.grid = nextGrid
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.backgroundColor)

	for x := 0; x < len(g.grid); x++ {
		for y := 0; y < len(g.grid[x]); y++ {
			if g.grid[x][y] == true {
				vector.FillRect(screen, float32(x*cellSize), float32(y*cellSize), cellSize, cellSize, color.White, true)
			}
		}
	}
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

			if g.grid[nx][ny] == true {
				aliveNeighbors++
			}
		}
	}

	if alive {
		return aliveNeighbors == 2 || aliveNeighbors == 3
	}
	return aliveNeighbors == 3
}
