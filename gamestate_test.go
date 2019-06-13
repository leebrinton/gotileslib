package tileslib

import "fmt"
import "testing"

func TestNewGameState(t *testing.T) {
	gs := NewGameState()

	if gs == nil {
		t.Error("Result of NewGameState() is nil!")
	}

	if gs.SolvedState != GAME_SOLVED {
		t.Error(fmt.Sprintf("SolvedState should be %d, but is %d", GAME_SOLVED, gs.SolvedState))
	}
}


func TestSetBitOn(t *testing.T) {
	gs := NewGameState()

	gs.SetBitOn(0)

	if gs.SolvedState != 1 {
		t.Error(fmt.Sprintf( "After SetBitOn(0) SolvedState should be 1 but is %d", gs.SolvedState))
	}
}

func TestSetBitOff(t *testing.T) {
	gs := NewGameState()

	gs.SetBitOn(0)
	gs.SetBitOff(0)

	if gs.SolvedState != GAME_SOLVED {
		t.Error(fmt.Sprintf( "After SetBitOff(0) SolvedState should be 0 but is %d", gs.SolvedState))
	}
}

func TestUpdate(t *testing.T) {
	gs := NewGameState()

	gs.Update(0, 14)	
	if gs.SolvedState != 1 {
		t.Error(fmt.Sprintf( "After Update(0, 14) SolvedState should be 1 but is %d", gs.SolvedState))
	}
	
	gs.Update(0, 0)
	if gs.SolvedState != GAME_SOLVED {
		t.Error(fmt.Sprintf( "After Update(0, 0 SolvedState should be 0 but is %d", gs.SolvedState))
	}
}

func TestSolved(t *testing.T) {
	gs := NewGameState()

	if !gs.Solved() {
		t.Error(fmt.Sprintf("GameState should start in the solved state but is %d", gs.SolvedState))
	}

	gs.SetBitOn(0)

	if gs.Solved() {
		t.Error(fmt.Sprintf("SolvedState should not be solved but has value %d", gs.SolvedState))
	}
}

func TestSolvedValue(t *testing.T) {
	gs := NewGameState()

	for i := 0; i < 8; i++ {
		gs.SetBitOn(i)
	}

	currVal := gs.SolvedValue()

	if currVal != 8 {
		t.Error(fmt.Sprintf("After setting the first 8 bits on there should still be 8 bits off giving a solved value of 8. SolvedValue = %d", currVal))
	}
}

func TestSolvedPercent(t *testing.T) {
	gs := NewGameState()

	for i := 0; i < 8; i++ {
		gs.SetBitOn(i)
	}

	currPct := gs.SolvedPercent()

	if currPct != 50 {
		t.Error(fmt.Sprintf("After setting the first 8 bits on there should still be 8 bits off giving a solved percent of 50. SolvedPercent = %d", currPct))
	}	
}

