package lexis

import (
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
	walkFrames     = 3
	animationSpeed = 8
	moveSpeed      = 3.0
)

type Player struct {
	sprite    *characterSprite
	currDir   Direction
	name      string
	x, y      float64
	isWalking bool
	animFrame int
	animTick  int
}

func newPlayer(sheet *ebiten.Image, startDir Direction, name string, x, y float64) *Player {
	return &Player{
		sprite:  newCharacterSprite(sheet),
		currDir: startDir,
		name:    name,
		x:       x,
		y:       y,
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

	length := math.Sqrt(dx*dx + dy*dy)
	p.x += (dx / length) * moveSpeed
	p.y += (dy / length) * moveSpeed
	p.isWalking = true
}

func (p *Player) update() {
	if p.isWalking {
		p.animTick++
		if p.animTick >= animationSpeed {
			p.animTick = 0
			p.animFrame = (p.animFrame + 1) % walkFrames
		}
	} else {
		p.animFrame = 0
		p.animTick = 0
	}
}

func (p *Player) currentImage() *ebiten.Image {
	return p.sprite.frame(p.currDir, p.isWalking, p.animFrame)
}
