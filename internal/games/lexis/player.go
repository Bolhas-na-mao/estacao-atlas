package lexis

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Direction int

const (
	Right Direction = iota
	Left
	Top
	Down
)

const (
	spriteWidth     = 92
	spriteHeight    = 92
	animationFrames = 6
	animationSpeed  = 6
	moveSpeed       = 1.5
)

type Player struct {
	idleSprites    map[Direction]*ebiten.Image
	walkingSprites map[Direction][]*ebiten.Image
	currDir        Direction
	name           string
	x, y           float64
	isWalking      bool
	animFrame      int
	animTick       int
}

func newPlayer(idleSpritesheet *ebiten.Image, walkingSpritesheets map[Direction]*ebiten.Image, startDir Direction, name string, x, y float64) *Player {
	idleSprites := map[Direction]*ebiten.Image{
		Right: idleSpritesheet.SubImage(image.Rect(0, 0, spriteWidth, spriteHeight)).(*ebiten.Image),
		Left:  idleSpritesheet.SubImage(image.Rect(spriteWidth, 0, spriteWidth*2, spriteHeight)).(*ebiten.Image),
	}

	walkingSprites := make(map[Direction][]*ebiten.Image)
	for dir, sheet := range walkingSpritesheets {
		frames := make([]*ebiten.Image, animationFrames)
		for i := 0; i < animationFrames; i++ {
			frames[i] = sheet.SubImage(image.Rect(i*spriteWidth, 0, (i+1)*spriteWidth, spriteHeight)).(*ebiten.Image)
		}
		walkingSprites[dir] = frames
	}

	return &Player{
		idleSprites:    idleSprites,
		walkingSprites: walkingSprites,
		currDir:        startDir,
		name:           name,
		x:              x,
		y:              y,
	}
}

func (p *Player) move(dx float64) {
	if dx == 0 {
		return
	}

	if dx < 0 {
		p.currDir = Left
	} else {
		p.currDir = Right
	}

	p.x += dx * moveSpeed
	p.isWalking = true
}

func (p *Player) update() {
	if p.isWalking {
		p.animTick++
		if p.animTick >= animationSpeed {
			p.animTick = 0
			p.animFrame = (p.animFrame + 1) % animationFrames
		}
	} else {
		p.animFrame = 0
		p.animTick = 0
	}
}

func (p *Player) currentImage() *ebiten.Image {
	if p.isWalking {
		return p.walkingSprites[p.currDir][p.animFrame]
	}
	return p.idleSprites[p.currDir]
}
