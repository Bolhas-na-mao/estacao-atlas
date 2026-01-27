package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const LAUNCHER_TITLE = "Estação Atlas"

func main() {
	game := NewGame()

	width, height := game.GetArea()

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(LAUNCHER_TITLE)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
