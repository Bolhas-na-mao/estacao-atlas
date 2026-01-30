package games

import (
	"errors"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/games/lexis"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ID     string
	Name   string
	Run    func(screen *ebiten.Image) error
	Update func() error
}

var games = []Game{
	{ID: "lexis", Name: "O Silêncio de Lexis", Run: lexis.Run, Update: lexis.Update},
}

var currentGame *Game

func ListGames() []Game {
	return games
}

func GetCurrentGame() *Game {
	return currentGame
}

func SetCurrentGame(id string) (*Game, error) {
	for i, g := range games {
		if g.ID == id {
			currentGame = &games[i]
			return &games[i], nil
		}
	}
	return nil, errors.New("game not found")
}

func UpdateCurrentGame() error {
	game := GetCurrentGame()

	if game == nil {
		return errors.New("no game selected")
	}

	if game.Update == nil {
		return errors.New("game has no update handler")
	}

	if err := game.Update(); err != nil {
		return err
	}

	return nil
}

func PlayCurrentGame(screen *ebiten.Image) error {

	game := GetCurrentGame()

	if game == nil {
		return errors.New("no game selected")
	}

	err := game.Run(screen)
	return err
}
