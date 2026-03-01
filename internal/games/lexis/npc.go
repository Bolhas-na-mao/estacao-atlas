package lexis

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	npcSpriteWidth  = 704 / 8
	npcSpriteHeight = 92
)

type Npc struct {
	idleSprites map[Direction]*ebiten.Image
	currDir     Direction
	name        string
	x, y        float64
	dialogue    []string
}

func newGolem(spritesheet *ebiten.Image, name string, x, y float64, dialogue []string) *Npc {
	south := spritesheet.SubImage(
		image.Rect(0, 0, npcSpriteWidth, npcSpriteHeight),
	).(*ebiten.Image)

	return &Npc{
		idleSprites: map[Direction]*ebiten.Image{Down: south},
		currDir:     Down,
		name:        name,
		x:           x,
		y:           y,
		dialogue:    dialogue,
	}
}

func (n *Npc) currentImage() *ebiten.Image {
	return n.idleSprites[n.currDir]
}
