package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

const SCREEN_WIDTH = 640
const SCREEN_HEIGHT = 480
const SQUARE_WIDTH = 30

func (g *Game) Update() error {
	return nil
}

func DrawGrid(screen *ebiten.Image) {

	for i := 0; i < SCREEN_WIDTH; i += SQUARE_WIDTH {
		for j := 0; j < SCREEN_HEIGHT; j += SQUARE_WIDTH {
			vector.StrokeRect(
				screen,
				float32(i), float32(j),
				SQUARE_WIDTH, SQUARE_WIDTH,
				1,
				color.RGBA{217, 218, 224, 255},
				false,
			)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{228, 228, 228, 255})
	DrawGrid(screen)

	ebitenutil.DebugPrint(screen, "Estação Atlas")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Estação Atlas")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
