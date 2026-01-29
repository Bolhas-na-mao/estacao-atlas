package main

import (
	"image/color"
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameState int

const (
	StateMenu GameState = iota
)

type Game struct {
	img          *ebiten.Image
	state        GameState
	screenWidth  int
	screenHeight int
}

const SCREEN_WIDTH = 1280
const SCREEN_HEIGHT = 720
const CELL_SIZE = 30
const PADDING = 10
const LOGO_PATH = "assets/atlas_logo.png"

func (g *Game) Update() error {
	return nil
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	screen.Fill(color.RGBA{228, 228, 228, 255})

	ui.DrawGrid(color.RGBA{217, 218, 224, 255}, screen, g.screenWidth, g.screenHeight, PADDING, CELL_SIZE)

	if g.img != nil {
		w, _ := g.img.Bounds().Dx(), g.img.Bounds().Dy()

		op := &ebiten.DrawImageOptions{}

		scale := 0.35
		op.GeoM.Scale(scale, scale)

		x := (float64(g.screenWidth) - (float64(w) * scale)) / 2
		y := float64(g.screenHeight) / 100

		op.GeoM.Translate(x, y)

		screen.DrawImage(g.img, op)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateMenu:
		g.drawMenu(screen)
	default:
		g.drawMenu(screen)
	}
}

func (g *Game) GetArea() (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.GetArea()
}

func NewGame() *Game {
	img, _, err := ebitenutil.NewImageFromFile(LOGO_PATH)
	if err != nil {
		log.Fatal(err)
	}

	return &Game{
		img:          img,
		state:        StateMenu,
		screenWidth:  SCREEN_WIDTH,
		screenHeight: SCREEN_HEIGHT,
	}
}
