package games

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type GameInfo struct {
	ID   string
	Name string
	New  func() Game
}

var registry []GameInfo

func Register(info GameInfo) {
	registry = append(registry, info)
}

func ListGames() []GameInfo {
	return registry
}
