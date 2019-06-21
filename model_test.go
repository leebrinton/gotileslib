package tileslib

import "fmt"
import "testing"

func TestNewCell(t *testing.T) {
	cell := NewCell()

	if cell == nil {
		t.Error("Result of NewCell() is nil!")
	}

}

func TestTransResult(t *testing.T) {
	tr := Pending

	if tr != 0 {
		t.Error(fmt.Sprintf("Var assigned as Pending should be 0 but is %d", tr))
	}

	tr = Ok

	if tr != 1 {
		t.Error(fmt.Sprintf("Var assigned as Ok should be 1 but is %d", tr))
	}

	tr = Exception

	if tr != 2 {
		t.Error(fmt.Sprintf("Var assigned as Exception should be 2 but is %d", tr))
	}

	tr = Error

	if tr != 3 {
		t.Error(fmt.Sprintf("Var assigned as Error should be 3 but is %d", tr))
	}
}

func TestModel(t *testing.T) {
	model := NewModel()

	if model.state.SolvedState != GAME_SOLVED {
		t.Error("SolvedState should start solved")
	}

	if model.emptyCellIndex != 0 {
		t.Error("emptyCellIndex should start at 0")
	}

	if model.lastTrans.sourceIndex != 0 {
		t.Error("lastTrans.sourceIndex should start at 0")
	}

	if model.lastTrans.destIndex != 0 {
		t.Error("lastTrans.destIndex should start at 0")
	}

	if model.StartTime.IsZero() {
		t.Error("Model.StartTime should have been initialized")
	}
}
