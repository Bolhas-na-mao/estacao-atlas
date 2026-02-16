package lexis

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Direction int

const (
	South Direction = iota
	North
	East
	West
)

const (
	spriteSize      = 48
	animationFrames = 4
	animationSpeed  = 8
	moveSpeed       = 3.0
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
		South: idleSpritesheet.SubImage(image.Rect(0, 0, spriteSize, spriteSize)).(*ebiten.Image),
		East:  idleSpritesheet.SubImage(image.Rect(spriteSize, 0, spriteSize*2, spriteSize)).(*ebiten.Image),
		North: idleSpritesheet.SubImage(image.Rect(spriteSize*2, 0, spriteSize*3, spriteSize)).(*ebiten.Image),
		West:  idleSpritesheet.SubImage(image.Rect(spriteSize*3, 0, spriteSize*4, spriteSize)).(*ebiten.Image),
	}

	walkingSprites := make(map[Direction][]*ebiten.Image)
	for dir, sheet := range walkingSpritesheets {
		frames := make([]*ebiten.Image, animationFrames)
		for i := 0; i < animationFrames; i++ {
			frames[i] = sheet.SubImage(image.Rect(i*spriteSize, 0, (i+1)*spriteSize, spriteSize)).(*ebiten.Image)
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

func (p *Player) move(dx, dy float64) {
	if dx == 0 && dy == 0 {
		return
	}

	if dx < 0 {
		p.currDir = West
	} else if dx > 0 {
		p.currDir = East
	} else if dy < 0 {
		p.currDir = North
	} else if dy > 0 {
		p.currDir = South
	}

	len := math.Sqrt(dx*dx + dy*dy)
	p.x += (dx / len) * moveSpeed
	p.y += (dy / len) * moveSpeed
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
