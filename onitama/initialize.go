package onitama

import "github.com/gorgonia/agogo/game"

func InitialState() OnitamaState {
	return OnitamaState{
		pawnBoard: []game.Colour{
			game.Black, game.Black, game.None, game.Black, game.Black,
			game.None, game.None, game.None, game.None, game.None,
			game.None, game.None, game.None, game.None, game.None,
			game.None, game.None, game.None, game.None, game.None,
			game.White, game.White, game.None, game.White, game.White,
		},
		kingBoard: []game.Colour{
			game.None, game.None, game.Black, game.None, game.None,
			game.None, game.None, game.None, game.None, game.None,
			game.None, game.None, game.None, game.None, game.None,
			game.None, game.None, game.None, game.None, game.None,
			game.None, game.None, game.White, game.None, game.None,
		},
		toMove:     game.Player(game.White),
		moveNumber: 0,
	}
}
