package lexis

import (
	"embed"
	"fmt"
	"image/color"
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var startButton *ui.Button
var hero *Character

//go:embed assets/*
var assets embed.FS

func init() {
	startButton = ui.NewButton(100, 100, 200, 50, "Lexis...")

	var err error
	hero, err = loadHeroAssets()
	if err != nil {
		log.Fatal(err)
	}
}

func loadHeroAssets() (*Character, error) {
	imgNorth, _, _ := ebitenutil.NewImageFromFile("internal/games/lexis/assets/characters/hero/idle/north.png")
	imgSouth, _, _ := ebitenutil.NewImageFromFile("internal/games/lexis/assets/characters/hero/idle/south.png")
	imgEast, _, _ := ebitenutil.NewImageFromFile("internal/games/lexis/assets/characters/hero/idle/east.png")
	imgWest, _, _ := ebitenutil.NewImageFromFile("internal/games/lexis/assets/characters/hero/idle/west.png")

	sprites := map[Direction]*ebiten.Image{
		North: imgNorth,
		South: imgSouth,
		East:  imgEast,
		West:  imgWest,
	}

	return NewCharacter(sprites, South, "Hero", 100, 100)
}

func Update() error {
	if startButton.Update() {
		fmt.Println("Botao clicado")
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		hero.Move(North)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		hero.Move(South)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		hero.Move(West)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		hero.Move(East)
	}

	return nil
}

func Run(screen *ebiten.Image) error {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	startButton.Draw(screen)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(hero.X, hero.Y)

	screen.DrawImage(hero.GetCurrImage(), op)

	return nil
}
