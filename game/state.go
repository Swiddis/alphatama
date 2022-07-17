package alphatama

import "github.com/gorgonia/agogo/game"

type OnitamaState struct {
	toMove     game.Player
	moveNumber int
}

func (s *OnitamaState) BoardSize() (int, int) {
	return 5, 5
}

func (s *OnitamaState) Board() []game.Colour {
	// TODO
	// 25 squares + 33 cards held by players + 33 neutral cards
	return make([]game.Colour, 91)
}

func (s *OnitamaState) ActionSpace() int {
	// 25 start squares * 25 end squares * 2 cards
	return 1250
}

func (s *OnitamaState) Hash() game.Zobrist {
	// TODO
	return 0
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
	return s == other.(*OnitamaState)
}

func (s *OnitamaState) Clone() game.State {
	return &OnitamaState{}
}
