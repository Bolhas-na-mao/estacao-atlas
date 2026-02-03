package lexis

import "github.com/hajimehoshi/ebiten/v2"

type Direction int

const (
	South Direction = iota
	North
	East
	West
)

type Character struct {
	Sprites map[Direction]*ebiten.Image
	CurrDir Direction
	Name    string
	X, Y    float64
}

func NewCharacter(sprites map[Direction]*ebiten.Image, startDir Direction, name string, X, Y float64) (*Character, error) {
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
