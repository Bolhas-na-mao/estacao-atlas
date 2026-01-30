package main

import (
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/launcher"
	"github.com/hajimehoshi/ebiten/v2"
)

const LAUNCHER_TITLE = "Estação Atlas"

func main() {
	launcher := launcher.NewLauncher()

	width, height := launcher.GetArea()

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(LAUNCHER_TITLE)

	if err := ebiten.RunGame(launcher); err != nil {
		log.Fatal(err)
	}
}
