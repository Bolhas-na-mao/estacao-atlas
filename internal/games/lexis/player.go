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
	moveSpeed      = 1.5
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

func (p *Player) move(dx, dy float64, isSolid func(x, y float64) bool) {
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
	vx := (dx / length) * moveSpeed
	vy := (dy / length) * moveSpeed

	if newX := p.x + vx; !p.hitsSolid(newX, p.y, isSolid) {
		p.x = newX
	}
	if newY := p.y + vy; !p.hitsSolid(p.x, newY, isSolid) {
		p.y = newY
	}

	p.isWalking = true
}

func (p *Player) hitsSolid(x, y float64, isSolid func(float64, float64) bool) bool {
	w := float64(spriteSize)
	margin := w / 5
	left := x + margin
	right := x + w - margin
	top := y + w/2
	bottom := y + w - 2

	return isSolid(left, top) || isSolid(right, top) ||
		isSolid(left, bottom) || isSolid(right, bottom)
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
