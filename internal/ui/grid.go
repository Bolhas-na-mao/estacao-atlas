package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawGrid(color color.RGBA, screen *ebiten.Image, screenWidth, screenHeight, padding, cellSize int) {

	if cellSize <= 0 || padding < 0 {
		return
	}

	availWidth := screenWidth - (padding * 2)
	availHeight := screenHeight - (padding * 2)

	cols := availWidth / cellSize
	rows := availHeight / cellSize

	totalGridWidth := cols * cellSize
	totalGridHeight := rows * cellSize

	startX := (screenWidth - totalGridWidth) / 2
	startY := (screenHeight - totalGridHeight) / 2

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {

			x := startX + (i * cellSize)
			y := startY + (j * cellSize)

			vector.StrokeRect(
				screen,
				float32(x), float32(y),
				float32(cellSize), float32(cellSize),
				2,
				color,
				false,
			)
		}
	}
}
