package onitama

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/gorgonia/agogo/game"
)

type CardMove struct {
	Dx int
	Dy int
}

type Card struct {
	Name  string     `json:"name"`
	Moves []CardMove `json:"moves"`
}

var zobristKeys [234]uint32
var Cards [34]Card

func loadKeys(seed int64) {
	rand.Seed(seed)
	if zobristKeys[0] == 0 {
		for i := 0; i < len(zobristKeys); i++ {
			zobristKeys[i] = rand.Uint32()
		}
	}
	rand.Seed(time.Now().UnixNano())
}

func loadCards(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Failed to load card file: " + err.Error())
	}
	err = json.Unmarshal(data, &Cards)
	if err != nil {
		panic("Failed to load card file: " + err.Error())
	}
}

func zobristHash(board []game.Colour) game.Zobrist {
	// For now just compute hashes the long way
	// TODO compute zobrist on the fly
	var hash uint32 = 0
	for i := 0; i < len(board); i++ {
		if board[i] == game.White {
			hash ^= zobristKeys[i]
		} else if board[i] == game.Black {
			hash ^= zobristKeys[116+i]
		}
	}
	return game.Zobrist(hash)
}

func initialPopulate(initial *OnitamaState) {
	blackPawns := []int{0, 1, 3, 4}
	whitePawns := []int{20, 21, 23, 24}
	for i := 0; i < 4; i++ {
		initial.pawnBoard[blackPawns[i]] = game.Black
		initial.pawnBoard[whitePawns[i]] = game.White
	}
	initial.kingBoard[2] = game.Black
	initial.kingBoard[22] = game.White
	cards := rand.Perm(34)
	initial.playerCards[cards[0]] = game.Black
	initial.playerCards[cards[1]] = game.Black
	initial.playerCards[cards[2]] = game.White
	initial.playerCards[cards[3]] = game.White
	initial.neutralCard = cards[4]
}

func InitialState() OnitamaState {
	initial := OnitamaState{
		pawnBoard:   make([]game.Colour, 25),
		kingBoard:   make([]game.Colour, 25),
		playerCards: make([]game.Colour, 34),
		neutralCard: 4,
		toMove:      game.Player(game.White),
		moveNumber:  0,
	}
	initialPopulate(&initial)
	initial.zobrist = zobristHash(initial.Board())
	return initial
}

func init() {
	loadKeys(-2542287859469082068)
	loadCards("cards.json")
}
