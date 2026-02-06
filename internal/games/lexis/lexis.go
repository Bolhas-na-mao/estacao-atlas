package lexis

import (
	"embed"
	"image/color"
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

var hero *Player

//go:embed assets/*
var assets embed.FS

func init() {
	idleSpritesheet, err := ui.RenderAsset(assets, "assets/hero/hero_idle.png")
	if err != nil {
		log.Fatal(err)
	}

	walkingSpritesheets := make(map[Direction]*ebiten.Image)

	southWalk, err := ui.RenderAsset(assets, "assets/hero/hero_walking_south.png")
	if err != nil {
		log.Fatal(err)
	}
	walkingSpritesheets[South] = southWalk

	northWalk, err := ui.RenderAsset(assets, "assets/hero/hero_walking_north.png")
	if err != nil {
		log.Fatal(err)
	}
	walkingSpritesheets[North] = northWalk

	eastWalk, err := ui.RenderAsset(assets, "assets/hero/hero_walking_east.png")
	if err != nil {
		log.Fatal(err)
	}
	walkingSpritesheets[East] = eastWalk

	westWalk, err := ui.RenderAsset(assets, "assets/hero/hero_walking_west.png")
	if err != nil {
		log.Fatal(err)
	}
	walkingSpritesheets[West] = westWalk

	hero, err = NewPlayer(idleSpritesheet, walkingSpritesheets, South, "Hero", 100, 100)
	if err != nil {
		log.Fatal(err)
	}
}

func Update() error {
	hero.IsWalking = false

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

	hero.Update()

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
