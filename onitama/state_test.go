package onitama

import (
	"testing"

	"github.com/gorgonia/agogo/game"
	"github.com/stretchr/testify/assert"
)

func TestOnitamaState_BoardSize(t *testing.T) {
	state := InitialState()
	sizeX, sizeY := state.BoardSize()
	assert.Equal(t, 5, sizeX)
	assert.Equal(t, 5, sizeY)
}

func TestOnitamaState_ToMove(t *testing.T) {
	state := InitialState()
	assert.Equal(t, state.ToMove(), game.Player(game.White))
}

func TestOnitamaState_SetToMove(t *testing.T) {
	state := InitialState()
	state.SetToMove(game.Player(game.Black))
	assert.Equal(t, state.ToMove(), game.Player(game.Black))
}

func TestOnitamaState_MoveNumber(t *testing.T) {
	state := InitialState()
	assert.Equal(t, 0, state.MoveNumber())
}

func TestOnitamaState_Passes(t *testing.T) {
	state := InitialState()
	assert.Equal(t, 0, state.Passes())
}

func TestOnitamaState_Handicap(t *testing.T) {
	state := InitialState()
	assert.Equal(t, 0, state.Handicap())
}

func TestOnitamaState_AdditionalScore(t *testing.T) {
	state := InitialState()
	assert.Equal(t, float32(0.0), state.AdditionalScore())
}

func TestOnitamaState_Eq(t *testing.T) {
	state1 := InitialState()
	state2 := InitialState()
	assert.True(t, state1.Eq(&state2))
}

func TestOnitamaState_Neq(t *testing.T) {
	state1 := InitialState()
	state2 := InitialState()
	state2.SetToMove(game.Player(game.Black))
	assert.False(t, state1.Eq(&state2))
}

func TestOnitamaState_Clone(t *testing.T) {
	state1 := InitialState()
	state2 := state1.Clone()
	state2.SetToMove(game.Player(game.Black))
	assert.False(t, state1.Eq(state2))
}
