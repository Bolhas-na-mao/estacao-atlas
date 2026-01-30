package games

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ID   string
	Name string
	Run  func(screen *ebiten.Image) error
}

func RunLexis(screen *ebiten.Image) error {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	return nil
}

var games = []Game{
	{ID: "lexis", Name: "O Silêncio de Lexis", Run: RunLexis},
}

var currentGame *Game

func ListGames() []Game {
	return games
}

func GetCurrentGame() *Game {
	return currentGame
}

func SetCurrentGame(id string) (*Game, error) {
	for _, g := range games {
		if g.ID == id {
			currentGame = &g
			return &g, nil
		}
	}
	return nil, errors.New("game not found")
}

func PlayCurrentGame(screen *ebiten.Image) error {
	game := GetCurrentGame()

	err := game.Run(screen)

	return err
}
