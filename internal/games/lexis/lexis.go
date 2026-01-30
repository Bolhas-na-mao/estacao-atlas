package lexis

import (
	"fmt"
	"image/color"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

var startButton *ui.Button

func init() {
	startButton = ui.NewButton(100, 100, 200, 50, "Lexis...")
}

func Update() error {
	if startButton.Update() {
		fmt.Println("Botao clicado")
	}

	return nil
}

func Run(screen *ebiten.Image) error {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	startButton.Draw(screen)

	return nil
}
