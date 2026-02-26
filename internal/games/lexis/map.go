package lexis

import "github.com/hajimehoshi/ebiten/v2"

type Room struct {
	sprite *ebiten.Image
	npcs   []Npc
}

type Map []Room
