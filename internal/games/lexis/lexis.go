package lexis

import (
	"embed"
	"image/color"
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/games"
	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

const heroScale = 3.0

func init() {
	games.Register(games.GameInfo{
		ID:   "lexis",
		Name: "O Silêncio de Lexis",
		New: func() games.Game {
			return New()
		},
	})
}

type LexisGame struct {
	hero *Player
}

func New() *LexisGame {
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

	hero := newPlayer(idleSpritesheet, walkingSpritesheets, South, "Hero", 100, 100)

	return &LexisGame{hero: hero}
}

func (g *LexisGame) Update() error {
	g.hero.isWalking = false

	var dx, dy float64
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		dy--
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		dy++
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx--
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx++
	}
	g.hero.move(dx, dy)

	g.hero.update()

	return nil
}

func (g *LexisGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(heroScale, heroScale)
	op.GeoM.Translate(g.hero.x, g.hero.y)

	screen.DrawImage(g.hero.currentImage(), op)
}
