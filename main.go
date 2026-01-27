package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	img *ebiten.Image
}

const SCREEN_WIDTH = 720
const SCREEN_HEIGHT = 480
const SQUARE_WIDTH = 30
const PADDING = 10

func (g *Game) Update() error {
	return nil
}

func DrawGrid(screen *ebiten.Image) {
	availWidth := SCREEN_WIDTH - (PADDING * 2)
	availHeight := SCREEN_HEIGHT - (PADDING * 2)

	cols := availWidth / SQUARE_WIDTH
	rows := availHeight / SQUARE_WIDTH

	totalGridWidth := cols * SQUARE_WIDTH
	totalGridHeight := rows * SQUARE_WIDTH

	startX := (SCREEN_WIDTH - totalGridWidth) / 2
	startY := (SCREEN_HEIGHT - totalGridHeight) / 2

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {

			x := startX + (i * SQUARE_WIDTH)
			y := startY + (j * SQUARE_WIDTH)

			vector.StrokeRect(
				screen,
				float32(x), float32(y),
				SQUARE_WIDTH, SQUARE_WIDTH,
				2,
				color.RGBA{217, 218, 224, 255},
				false,
			)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{228, 228, 228, 255})
	DrawGrid(screen)

	if g.img != nil {
		w, _ := g.img.Bounds().Dx(), g.img.Bounds().Dy()

		op := &ebiten.DrawImageOptions{}

		scale := 0.35
		op.GeoM.Scale(scale, scale)

		x := (float64(SCREEN_WIDTH) - (float64(w) * scale)) / 2
		y := float64(SCREEN_HEIGHT) / 100

		op.GeoM.Translate(x, y)

		screen.DrawImage(g.img, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	g := &Game{}

	var err error
	g.img, _, err = ebitenutil.NewImageFromFile("assets/atlas_logo.png")

	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Estação Atlas")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
