package lexis

import (
	"embed"
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/games"
	"github.com/Bolhas-na-mao/estacao-atlas/internal/logger"
	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

const (
	heroScale = 2.5
	screenW   = 1280
	screenH   = 720
)

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
	hero     *Player
	worldMap *WorldMap
	camera   *Camera
}

func New() *LexisGame {
	logger.Info("starting Lexis")
	heroSheet, err := ui.RenderAsset(assets, "assets/characters/hero.png")
	if err != nil {
		log.Fatal(err)
	}
	project, err := parseLdtk(assets, "assets/lexis.ldtk")
	if err != nil {
		log.Fatal(err)
	}

	worldMap := newWorldMap(project, assets)

	hero := newPlayer(heroSheet, South, 100, 160)

	room0 := worldMap.current()
	heroCenterX := hero.x + float64(spriteSize)/2
	heroCenterY := hero.y + float64(spriteSize)/2
	camera := newCamera(heroCenterX, heroCenterY, room0.width, room0.height, screenW, screenH)

	return &LexisGame{hero: hero, worldMap: worldMap, camera: camera}
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

	g.hero.move(dx, dy, g.worldMap.current().isSolid)

	room := g.worldMap.current()
	g.hero.x = clamp(g.hero.x, 0, float64(room.width-spriteSize))
	g.hero.y = clamp(g.hero.y, 0, float64(room.height-spriteSize))

	heroCenterX := g.hero.x + float64(spriteSize)/2
	heroCenterY := g.hero.y + float64(spriteSize)/2
	cur := g.worldMap.current()
	g.camera.update(heroCenterX, heroCenterY, cur.width, cur.height)

	g.hero.update()

	return nil
}

func (g *LexisGame) Draw(screen *ebiten.Image) {
	g.worldMap.draw(screen, g.camera)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(heroScale, heroScale)
	op.GeoM.Translate(
		(g.hero.x-g.camera.x)*heroScale,
		(g.hero.y-g.camera.y)*heroScale,
	)
	screen.DrawImage(g.hero.currentImage(), op)
}
