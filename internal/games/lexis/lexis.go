package lexis

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func Run(screen *ebiten.Image) error {
	screen.Fill(color.RGBA{228, 228, 228, 255})

	return nil
}
