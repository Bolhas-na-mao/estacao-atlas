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
	hero   *Player
	theMap Map
	camera Camera
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

	hero := newPlayer(idleSpritesheet, walkingSpritesheets, Right, "Hero", 30, 30)

	lib1, err := ui.RenderAsset(assets, "assets/rooms/library.png")
	if err != nil {
		log.Fatal(err)
	}
	lib2, err := ui.RenderAsset(assets, "assets/rooms/hall.png")
	if err != nil {
		log.Fatal(err)
	}

	golemSheet, err := ui.RenderAsset(assets, "assets/npcs/golem.png")
	if err != nil {
		log.Fatal(err)
	}
	golem := newGolem(golemSheet, "Golem", 390, 30)

	theMap := Map{
		rooms: []Room{
			{sprite: lib1, worldX: 0},
			{sprite: lib2, worldX: roomWidth, npcs: []Npc{*golem}},
		},
	}

	return &LexisGame{
		hero:   hero,
		theMap: theMap,
		camera: Camera{},
	}
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

	if g.hero.x < 0 {
		g.hero.x = 0
	}
	if g.hero.x > worldWidth-spriteWidth {
		g.hero.x = worldWidth - spriteWidth
	}

	g.hero.update()
	g.camera.update(g.hero.x, worldWidth)

	return nil
}

func (g *LexisGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	g.theMap.draw(screen, &g.camera)

	sx, sy := g.camera.toScreen(g.hero.x, g.hero.y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(heroScale, heroScale)
	op.GeoM.Translate(sx, sy)
	screen.DrawImage(g.hero.currentImage(), op)
}
