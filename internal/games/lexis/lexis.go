package lexis

import (
	"embed"
	"fmt"
	"image/color"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

var startButton *ui.Button

//go:embed assets/*
var assets embed.FS

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

	img, err := ui.RenderAsset(assets, "assets/characters/hero/idle/east.png")

	if err != nil {
		return err
	}

	screen.DrawImage(img, nil)

	return nil
}
