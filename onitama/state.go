package onitama

import (
	"reflect"

	"github.com/gorgonia/agogo/game"
)

type OnitamaState struct {
	pawnBoard   []game.Colour
	kingBoard   []game.Colour
	playerCards []game.Colour
	neutralCard int
	toMove      game.Player
	moveNumber  int
	zobrist     game.Zobrist
}

func (s *OnitamaState) BoardSize() (int, int) {
	return 5, 5
}

func (s *OnitamaState) Board() []game.Colour {
	// 25 pawn squares + 25 king squares + 33 cards held by players + 33 neutral cards
	board := make([]game.Colour, 116)
	for i := 0; i < 25; i++ {
		board[i] = s.pawnBoard[i]
		board[i+25] = s.kingBoard[i]
	}
	for i := 0; i < 33; i++ {
		board[i+50] = s.playerCards[i]
	}
	board[83+s.neutralCard] = game.White
	return board
}

func (s *OnitamaState) ActionSpace() int {
	// 25 start squares * 25 end squares * 2 cards
	return 1250
}

func (s *OnitamaState) Hash() game.Zobrist {
	return s.zobrist
}

func (s *OnitamaState) ToMove() game.Player {
	return s.toMove
}

func (s *OnitamaState) Passes() int {
	return 0
}

func (s *OnitamaState) MoveNumber() int {
	return s.moveNumber
}

func (s *OnitamaState) LastMove() game.PlayerMove {
	// TODO
	return game.PlayerMove{}
}

func (s *OnitamaState) Handicap() int {
	return 0
}

func (s *OnitamaState) Score(player game.Player) float32 {
	// TODO
	return 0.0
}

func (s *OnitamaState) AdditionalScore() float32 {
	return 0.0
}

func (s *OnitamaState) Ended() (bool, game.Player) {
	// TODO
	return false, game.Player(game.None)
}

func (s *OnitamaState) SetToMove(player game.Player) {
	s.toMove = player
}

func (s *OnitamaState) Check(move game.PlayerMove) bool {
	// TODO
	return false
}

func (s *OnitamaState) Apply(move game.PlayerMove) game.State {
	// TODO
	return &OnitamaState{}
}

func (s *OnitamaState) Reset() {
	// TODO
}

func (s *OnitamaState) Historical(i int) []game.Colour {
	// TODO
	return []game.Colour{}
}

func (s *OnitamaState) UndoLastMove() {
	// TODO
}

func (s *OnitamaState) Fwd() {
	// TODO
}

func (s *OnitamaState) Eq(other game.State) bool {
	return reflect.DeepEqual(s, other.(*OnitamaState))
}

func (s *OnitamaState) Clone() game.State {
	return &OnitamaState{}
}
