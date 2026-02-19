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
	idleSpritesheet, err := ui.RenderAsset(assets, "assets/hero/hero.png")
	if err != nil {
		log.Fatal(err)
	}

	walkingSpritesheets := make(map[Direction]*ebiten.Image)

	rightWalk, err := ui.RenderAsset(assets, "assets/hero/walking_right.png")
	if err != nil {
		log.Fatal(err)
	}
	walkingSpritesheets[Right] = rightWalk

	leftWalk, err := ui.RenderAsset(assets, "assets/hero/walking_left.png")
	if err != nil {
		log.Fatal(err)
	}
	walkingSpritesheets[Left] = leftWalk

	hero := newPlayer(idleSpritesheet, walkingSpritesheets, Right, "Hero", 100, 100)

	return &LexisGame{hero: hero}
}

func (g *LexisGame) Update() error {
	g.hero.isWalking = false

	var dx float64
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx--
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx++
	}
	g.hero.move(dx)

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
