package lexis

import "github.com/hajimehoshi/ebiten/v2"

type Npc struct {
	idleSprites map[Direction]*ebiten.Image
	currDir     Direction
	name        string
	x, y        float64
}
