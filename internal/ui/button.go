package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	X, Y, W, H int
	Text       string
	IsHovered  bool
}

func NewButton(x, y, w, h int, text string) *Button {
	return &Button{
		X:    x,
		Y:    y,
		W:    w,
		H:    h,
		Text: text,
	}
}

func (b *Button) Update() bool {
	mx, my := ebiten.CursorPosition()

	if mx >= b.X && mx <= b.X+b.W && my >= b.Y && my <= b.Y+b.H {
		b.IsHovered = true

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			return true
		}
	} else {
		b.IsHovered = false
	}

	return false
}

func (b *Button) Draw(screen *ebiten.Image) {
	c := color.RGBA{100, 100, 100, 255}
	if b.IsHovered {
		c = color.RGBA{150, 150, 150, 255}
	}

	vector.DrawFilledRect(screen, float32(b.X), float32(b.Y), float32(b.W), float32(b.H), c, false)

	ebitenutil.DebugPrintAt(screen, b.Text, b.X+10, b.Y+10)
}
