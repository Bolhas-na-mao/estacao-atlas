package lexis

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

var hero *Character

//go:embed assets/*
var assets embed.FS

func init() {
	var err error
	hero, err = loadHeroAssets()
	if err != nil {
		log.Fatal(err)
	}
}

func loadHeroAssets() (*Character, error) {
	spritesheet, err := ui.RenderAsset(assets, "assets/hero/hero_idle.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load hero spritesheet: %w", err)
	}

	spriteWidth := 48

	sprites := map[Direction]*ebiten.Image{
		South: spritesheet.SubImage(image.Rect(0, 0, spriteWidth, spriteWidth)).(*ebiten.Image),
		East:  spritesheet.SubImage(image.Rect(spriteWidth, 0, spriteWidth*2, spriteWidth)).(*ebiten.Image),
		North: spritesheet.SubImage(image.Rect(spriteWidth*2, 0, spriteWidth*3, spriteWidth)).(*ebiten.Image),
		West:  spritesheet.SubImage(image.Rect(spriteWidth*3, 0, spriteWidth*4, spriteWidth)).(*ebiten.Image),
	}

	return NewCharacter(sprites, South, "Hero", 100, 100)
}

func Update() error {
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

	op := &ebiten.DrawImageOptions{}

	scale := 3.0

	op.GeoM.Scale(scale, scale)

	op.GeoM.Translate(hero.X, hero.Y)

	screen.DrawImage(hero.GetCurrImage(), op)

	return nil
}
