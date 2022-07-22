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
	prev        *OnitamaState
	next        *OnitamaState
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
	for i := 0; i < 34; i++ {
		board[i+50] = s.playerCards[i]
	}
	board[84+s.neutralCard] = game.White
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
	if s.kingBoard[2] == game.White {
		return true, game.Player(game.White)
	} else if s.kingBoard[22] == game.Black {
		return true, game.Player(game.Black)
	}
	var whiteKing, blackKing bool
	for i := 0; i < 25; i++ {
		if s.kingBoard[i] == game.White {
			whiteKing = true
		} else if s.kingBoard[i] == game.Black {
			blackKing = true
		}
	}
	if !whiteKing {
		return true, game.Player(game.Black)
	} else if !blackKing {
		return true, game.Player(game.White)
	}
	return false, game.Player(game.None)
}

func (s *OnitamaState) SetToMove(player game.Player) {
	s.toMove = player
}

func getMoveData(move game.PlayerMove) (card, start, end int) {
	if move.Single > 625 {
		card = 1
	}
	start = (int(move.Single) / 25) % 25
	end = int(move.Single) % 25
	return card, start, end
}

func checkMoveCard(cardIdx, start, end int, color game.Colour) bool {
	sx, sy := start%25, start/25
	ex, ey := end%25, end/25
	dx := ex - sx
	dy := sy - ey
	if color == game.White {
		dx *= -1
		dy *= -1
	}
	for i := 0; i < len(cards[cardIdx].moves); i++ {
		if cards[cardIdx].moves[i].dx == dx && cards[cardIdx].moves[i].dy == dy {
			return true
		}
	}
	return false
}

func (s *OnitamaState) Check(move game.PlayerMove) bool {
	card, start, end := getMoveData(move)

	// Can only start from one of our pieces
	if s.pawnBoard[start] != game.Colour(move.Player) && s.kingBoard[start] != game.Colour(move.Player) {
		return false
	}
	// Avoid self-captures
	if s.pawnBoard[end] == game.Colour(move.Player) || s.kingBoard[end] == game.Colour(move.Player) {
		return false
	}

	// Get which card the move is using
	var cidx, count int
	for i := 0; i < 33; i++ {
		if s.playerCards[i] == game.Colour(move.Player) && count <= card {
			cidx = i
			count++
		}
	}
	// Check that the move is on the card
	return checkMoveCard(cidx, start, end, game.Colour(move.Player))
}

func (s *OnitamaState) Apply(move game.PlayerMove) game.State {
	card, start, end := getMoveData(move)
	next := &OnitamaState{}

	// Metadata
	if s.toMove == game.Player(game.White) {
		next.toMove = game.Player(game.Black)
	} else {
		next.toMove = game.Player(game.White)
	}
	next.moveNumber = s.moveNumber + 1
	next.zobrist = s.zobrist
	next.prev = s
	s.next = next

	// Moving pieces
	next.pawnBoard = make([]game.Colour, 25)
	next.kingBoard = make([]game.Colour, 25)
	for i := 0; i < 25; i++ {
		next.pawnBoard[i] = s.pawnBoard[i]
		next.kingBoard[i] = s.kingBoard[i]
	}
	next.pawnBoard[end] = s.pawnBoard[start]
	s.pawnBoard[start] = game.None
	next.kingBoard[end] = s.kingBoard[start]
	s.kingBoard[start] = game.None

	// Swapping cards
	var cidx, count int
	next.playerCards = make([]game.Colour, 33)
	for i := 0; i < 34; i++ {
		next.playerCards[i] = s.playerCards[i]
		if s.playerCards[i] == game.Colour(move.Player) && count <= card {
			cidx = i
			count++
		}
	}
	next.playerCards[cidx] = game.None
	next.neutralCard = cidx
	next.playerCards[s.neutralCard] = game.Colour(move.Player)

	// TODO compute zobrist intelligently
	next.zobrist = zobristHash(next.Board())
	return next
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
