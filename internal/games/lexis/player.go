package lexis

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Direction int

const (
	South Direction = iota
	North
	East
	West
)

const SPRITE_SIZE = 48
const ANIMATION_FRAMES = 4
const ANIMATION_SPEED = 8

type Player struct {
	IdleSprites    map[Direction]*ebiten.Image
	WalkingSprites map[Direction][]*ebiten.Image
	CurrDir        Direction
	Name           string
	X, Y           float64
	IsWalking      bool
	AnimFrame      int
	AnimTick       int
}

func NewPlayer(idleSpritesheet *ebiten.Image, walkingSpritesheets map[Direction]*ebiten.Image, startDir Direction, name string, X, Y float64) (*Player, error) {
	idleSprites := map[Direction]*ebiten.Image{
		South: idleSpritesheet.SubImage(image.Rect(0, 0, SPRITE_SIZE, SPRITE_SIZE)).(*ebiten.Image),
		East:  idleSpritesheet.SubImage(image.Rect(SPRITE_SIZE, 0, SPRITE_SIZE*2, SPRITE_SIZE)).(*ebiten.Image),
		North: idleSpritesheet.SubImage(image.Rect(SPRITE_SIZE*2, 0, SPRITE_SIZE*3, SPRITE_SIZE)).(*ebiten.Image),
		West:  idleSpritesheet.SubImage(image.Rect(SPRITE_SIZE*3, 0, SPRITE_SIZE*4, SPRITE_SIZE)).(*ebiten.Image),
	}

	walkingSprites := make(map[Direction][]*ebiten.Image)
	for dir, sheet := range walkingSpritesheets {
		frames := make([]*ebiten.Image, ANIMATION_FRAMES)
		for i := 0; i < ANIMATION_FRAMES; i++ {
			frames[i] = sheet.SubImage(image.Rect(i*SPRITE_SIZE, 0, (i+1)*SPRITE_SIZE, SPRITE_SIZE)).(*ebiten.Image)
		}
		walkingSprites[dir] = frames
	}

	return &Player{
		IdleSprites:    idleSprites,
		WalkingSprites: walkingSprites,
		CurrDir:        startDir,
		Name:           name,
		X:              X,
		Y:              Y,
		IsWalking:      false,
		AnimFrame:      0,
		AnimTick:       0,
	}, nil
}

func (c *Player) Move(dir Direction) {
	c.CurrDir = dir
	c.IsWalking = true

	speed := 3.0
	switch dir {
	case North:
		c.Y -= speed
	case South:
		c.Y += speed
	case West:
		c.X -= speed
	case East:
		c.X += speed
	}
}

func (c *Player) Update() {
	if c.IsWalking {
		c.AnimTick++
		if c.AnimTick >= ANIMATION_SPEED {
			c.AnimTick = 0
			c.AnimFrame = (c.AnimFrame + 1) % ANIMATION_FRAMES
		}
	} else {
		c.AnimFrame = 0
		c.AnimTick = 0
	}
}

func (c *Player) GetCurrImage() *ebiten.Image {
	if c.IsWalking {
		return c.WalkingSprites[c.CurrDir][c.AnimFrame]
	}
	return c.IdleSprites[c.CurrDir]
}
