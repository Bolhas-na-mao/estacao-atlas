package games

import "errors"

type Game struct {
	ID   string
	Name string
}

var games = []Game{
	{ID: "lexis", Name: "O Silêncio de Lexis"},
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
