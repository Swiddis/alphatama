package main

import (
	"log"
	"time"

	"github.com/Swiddis/alphatama/onitama"
	"github.com/gorgonia/agogo"
	dual "github.com/gorgonia/agogo/dualnet"
	"github.com/gorgonia/agogo/game"
	"github.com/gorgonia/agogo/mcts"
)

// Not really traditional FEN but whatever
func fen(s game.State) string {
	state := s.(*onitama.OnitamaState)
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
	conf := agogo.Config{
		Name:            "Onitama",
		NNConf:          dual.DefaultConf(5, 5, 1250),
		MCTSConf:        mcts.DefaultConfig(3),
		UpdateThreshold: 0.52,
	}
	conf.NNConf.BatchSize = 100
	conf.NNConf.Features = 2
	conf.NNConf.K = 3
	conf.NNConf.SharedLayers = 3
	conf.MCTSConf = mcts.Config{
		PUCT:           1.0,
		M:              3,
		N:              3,
		Timeout:        100 * time.Millisecond,
		PassPreference: mcts.DontPreferPass,
		Budget:         1000,
		DumbPass:       true,
		RandomCount:    0,
	}
	// conf.Encoder = fen
	g := onitama.InitialState()
	agg := agogo.New(&g, conf)
	err := agg.Learn(5, 50, 100, 100)
	if err != nil {
		log.Println(err)
	}
	agg.Save("example.model")

	// // var initial *onitama.OnitamaState
	// initial := onitama.InitialState()
	// state := &initial
	// // fmt.Printf("%+v\n", initial)
	// count := 0
	// end, _ := state.Ended()
	// for !end {
	// 	fmt.Println(fen(state))
	// 	perm := rand.Perm(1250)
	// 	for i := 0; i < 1250; i++ {
	// 		move := game.PlayerMove{
	// 			Single: game.Single(perm[i]),
	// 			Player: state.ToMove(),
	// 		}
	// 		if state.Check(move) {
	// 			state = state.Apply(move).(*onitama.OnitamaState)
	// 			break
	// 		} else if i == 1249 {
	// 			panic("no legal moves")
	// 		}
	// 	}
	// 	count++
	// 	if count > 500 {
	// 		break
	// 	}
	// 	end, _ = state.Ended()
	// }
	// fmt.Println(fen(state))
}
