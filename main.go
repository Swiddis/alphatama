package main

import (
	"fmt"

	"github.com/Swiddis/alphatama/onitama"
	"github.com/gorgonia/agogo/game"
)

// Not really traditional FEN but whatever
func fen(state onitama.OnitamaState) string {
	fenstr := ""
	board := state.Board()
	for i := 0; i < 25; i++ {
		if board[i] == game.White {
			fenstr += "P"
		} else if board[i] == game.Black {
			fenstr += "p"
		} else if board[i+25] == game.White {
			fenstr += "K"
		} else if board[i+25] == game.Black {
			fenstr += "k"
		} else {
			fenstr += "."
		}
		if i%5 == 4 && i < 24 {
			fenstr += "/"
		}
	}
	fenstr += " "
	if game.Colour(state.ToMove()) == game.White {
		fenstr += "w"
	} else {
		fenstr += "b"
	}
	fenstr += " "
	for i := 0; i < 34; i++ {
		if board[i+50] == game.White {
			fenstr += onitama.Cards[i].Name + ","
		}
	}
	fenstr = fenstr[:len(fenstr)-1] + " "
	for i := 0; i < 34; i++ {
		if board[i+50] == game.Black {
			fenstr += onitama.Cards[i].Name + ","
		}
	}
	fenstr = fenstr[:len(fenstr)-1] + " "
	for i := 0; i < 34; i++ {
		if board[i+84] == game.White {
			fenstr += onitama.Cards[i].Name
		}
	}
	return fenstr
}

func main() {
	initial := onitama.InitialState()
	// fmt.Printf("%+v\n", initial)
	fmt.Println(fen(initial))
}
