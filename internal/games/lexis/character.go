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

type Character struct {
	Sprites map[Direction]*ebiten.Image
	CurrDir Direction
	Name    string
	X, Y    float64
}

func NewCharacter(spritesheet *ebiten.Image, startDir Direction, name string, X, Y float64) (*Character, error) {

	sprites := map[Direction]*ebiten.Image{
		South: spritesheet.SubImage(image.Rect(0, 0, SPRITE_SIZE, SPRITE_SIZE)).(*ebiten.Image),
		East:  spritesheet.SubImage(image.Rect(SPRITE_SIZE, 0, SPRITE_SIZE*2, SPRITE_SIZE)).(*ebiten.Image),
		North: spritesheet.SubImage(image.Rect(SPRITE_SIZE*2, 0, SPRITE_SIZE*3, SPRITE_SIZE)).(*ebiten.Image),
		West:  spritesheet.SubImage(image.Rect(SPRITE_SIZE*3, 0, SPRITE_SIZE*4, SPRITE_SIZE)).(*ebiten.Image),
	}

	return &Character{
		Sprites: sprites,
		CurrDir: startDir,
		Name:    name,
		X:       X,
		Y:       Y,
	}, nil
}

func (c *Character) Move(dir Direction) {
	c.CurrDir = dir

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

func (c *Character) GetCurrImage() *ebiten.Image {
	return c.Sprites[c.CurrDir]
}
