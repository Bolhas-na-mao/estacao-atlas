package lexis

type Direction int

const (
	South Direction = iota
	North
	East
	West
)

type Sprite struct {
	path      string
	direction Direction
}

type Character struct {
	Sprites []Sprite
	CurrDir Direction
	Name    string
}

func NewCharacter(sprites []Sprite, startDir Direction, name string) (Character, error) {

	// check stuff later

	return Character{
		Sprites: sprites,
		CurrDir: startDir,
		Name:    name,
	}, nil
}
