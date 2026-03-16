package lexis

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const spriteSize = 32

type characterSprite struct {
	idle    map[Direction]*ebiten.Image
	walking map[Direction][]*ebiten.Image
}

func newCharacterSprite(sheet *ebiten.Image) *characterSprite {
	dirOffset := map[Direction]int{
		South: 0,
		North: 4,
		East:  8,
		West:  12,
	}

	idle := make(map[Direction]*ebiten.Image, 4)
	walking := make(map[Direction][]*ebiten.Image, 4)

	for dir, col := range dirOffset {
		x := col * spriteSize
		idle[dir] = sheet.SubImage(image.Rect(x, 0, x+spriteSize, spriteSize)).(*ebiten.Image)

		frames := make([]*ebiten.Image, walkFrames)
		for i := range frames {
			fx := (col + 1 + i) * spriteSize
			frames[i] = sheet.SubImage(image.Rect(fx, 0, fx+spriteSize, spriteSize)).(*ebiten.Image)
		}
		walking[dir] = frames
	}

	return &characterSprite{idle: idle, walking: walking}
}

func (cs *characterSprite) frame(dir Direction, walking bool, animFrame int) *ebiten.Image {
	if walking {
		return cs.walking[dir][animFrame]
	}
	return cs.idle[dir]
}
