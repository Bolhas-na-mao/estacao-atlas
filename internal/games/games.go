package games

import (
	"errors"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/games/lexis"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ID   string
	Name string
	Run  func(screen *ebiten.Image) error
}

var games = []Game{
	{ID: "lexis", Name: "O Silêncio de Lexis", Run: lexis.Run},
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

	if game == nil {
		return errors.New("no game selected")
	}

	err := game.Run(screen)
	return err
}
