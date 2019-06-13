package tileslib

import "fmt"
import "time"

const MAX_GAME_CELL_INDEX = byte(15)
const NUM_GAME_COLS = byte(4)
const NUM_GAME_ROWS = byte(4)
const CELL_NOT_FOUND = -1

type Cell struct {
	cindex byte
	value  byte
}

func NewCell() *Cell {
	cell := new(Cell)

	return cell
}

func (cell Cell) String() string {
	return fmt.Sprintf("%2d", cell.value)
}

type TransResult int

const (
	Pending TransResult = iota
	Ok
	Exception
	Error
)

func (tr TransResult) String() string {
	result := "unknown"

	switch tr {
	case Pending:
		result = "Pending"

	case Ok:
		result = "Ok"

	case Exception:
		result = "Exception"

	case Error:
		result = "Error"
	}
	return result
}

type GameTransaction struct {
	sourceIndex byte
	destIndex   byte
	result      TransResult
}

func (gt GameTransaction) Markdown() string {
	result := fmt.Sprintf("* sourceIndex [%d]\n", gt.sourceIndex)
	result += fmt.Sprintf("* destIndex [%d]\n", gt.destIndex)
	result += fmt.Sprintf("* result [%s]\n", gt.result.String())

	return result
}

func (gt GameTransaction) String() string {
	str := fmt.Sprintf("srcIndex = [%d]\n", gt.sourceIndex)
	str += fmt.Sprintf("destIndex = [%d]\n", gt.destIndex)
	str += fmt.Sprintf("result = [%s]\n", gt.result.String())

	return str
}

type Model struct {
	state          *GameState
	StartTime      time.Time
	emptyCellIndex byte
	cells          [NUM_GAME_CELLS]Cell
	lastTrans      GameTransaction
}

func NewModel() *Model {
	m := new(Model)

	m.state = NewGameState()
	m.StartTime = time.Now()
	
	return m
}

func (m *Model) SolvedValue() int {
	return m.state.SolvedValue() 
}

func (m *Model) SolvedPercent() int {
	return m.state.SolvedPercent()
}

func (m *Model) Solved() bool {
	return m.state.Solved()
}

func (m *Model) StartNewGame( scrambleIterations int) {
	if scrambleIterations == 0 {
		scrambleIterations = DEFAULT_SCRAMBLE_ITERATIONS;
	}

	m.state.SolvedState = GAME_SOLVED;
	m.loadCells();
	m.scramble( scrambleIterations );
	m.StartTime = time.Now();
}

func (m *Model) CellValueAt(index int) byte {
	return m.cells[index].value
}

func (m *Model) LastSource() byte {
	return m.lastTrans.sourceIndex
}

func (m *Model) LastDest() byte {
	return m.lastTrans.destIndex
}

func (m *Model) LastTransResult() TransResult {
	return m.lastTrans.result
}

func (m *Model) updateState(index int) {
	ci := m.cells[index].cindex

	m.state.Update(index, int(ci))
}

func (m *Model) loadCells() {
	for i := 0; i < NUM_GAME_CELLS; i++ {
		m.cells[i].cindex = byte(i)
		m.cells[i].value = byte(i + 1)
	}
	m.emptyCellIndex = MAX_GAME_CELL_INDEX
}

func (m *Model)scramble(iterations int) {
	for i := 0; i < iterations; i++ {
		d := RandomDirection()
		m.moveEmptyCell(d)
	}
}

func (m *Model) swapCells(index1 int, index2 int) int {
	if index1 == CELL_NOT_FOUND || index2 == CELL_NOT_FOUND {
		return CELL_NOT_FOUND;
	}
	
	var temp Cell
	temp.cindex = m.cells[index1].cindex;
	temp.value = m.cells[index1].value;
            
	m.cells[index1].cindex = m.cells[index2].cindex;
	m.cells[index1].value = m.cells[index2].value;

	m.cells[index2].cindex = temp.cindex;
	m.cells[index2].value = temp.value;

	m.updateState(index1)
	m.updateState(index2)

	return index2
}

func RowFromIndex(index byte) byte {
	return byte(index / NUM_GAME_ROWS)
}

func ColFromIndex(index byte) byte {
	return (byte) (index % NUM_GAME_COLS);
}

func indexOfCellAbove(cindex byte) int {
	result := CELL_NOT_FOUND
	row := RowFromIndex(cindex)

	if row > 0 {
		result = int(cindex - NUM_GAME_COLS)
	}
	return result
}

func indexOfCellBelow(cindex byte) int {
	result := CELL_NOT_FOUND
	row := RowFromIndex(cindex)

	if row < 3 {
		result = int(cindex + NUM_GAME_COLS)
	}
	return result
}

func indexOfCellLeft(cindex byte) int {
	result := CELL_NOT_FOUND
	col := ColFromIndex(cindex)

	if col > 0 {
		result = (int(cindex) - 1)
	}
	return result
}

func indexOfCellRight(cindex byte) int {
	result := CELL_NOT_FOUND
	col := ColFromIndex(cindex)

	if col < 3 {
		result = (int(cindex) + 1)
	}
	return result
}

func (m *Model) getDestIndex(direction Direction) int {
	destIndex := 0

	switch direction {
	case Up:
		destIndex = indexOfCellAbove(m.emptyCellIndex)

	case Down:
		destIndex = indexOfCellBelow(m.emptyCellIndex)

	case Left:
		destIndex = indexOfCellLeft(m.emptyCellIndex)

	case Right:
		destIndex = indexOfCellRight(m.emptyCellIndex)

	}
	return destIndex
}

func (m *Model) moveEmptyCell(direction Direction) {
	destIndex := m.getDestIndex(direction)

	m.lastTrans.sourceIndex = byte(m.emptyCellIndex)
	m.lastTrans.destIndex = byte(destIndex)
	m.lastTrans.result = Pending
	
	swapResult := m.swapCells(int(m.emptyCellIndex), destIndex)

	if swapResult == CELL_NOT_FOUND {
		m.lastTrans.result = Exception
	} else {
		m.lastTrans.result = Ok
		m.emptyCellIndex = byte(destIndex)
	}
}

func (m *Model) moveValueCell(directions Direction) {
	m.moveEmptyCell( ReverseDirection(directions))
}

func (m *Model) MoveCell(directions Direction, mode CommandModeType) {
	if mode == EmptyCellCentric {
		m.moveEmptyCell(directions)
	} else {
		m.moveValueCell(directions)
	}
}

func cellArrayToHtml(array *[NUM_GAME_CELLS]Cell) string {
	result := "<table><tr>"

	for i:= 0; i < NUM_GAME_CELLS; i++ {
		result += "<td>"
		result += fmt.Sprintf("%d", array[i].value)
		result += "</td>"
	}
	result += "</tr></table>"
	return result
}

func (m *Model) Markdown() string {
	result := "### state ###\n"
	result += m.state.Markdown()
	result += "\n"

	result += "### StartTime ###\n"
	result += m.StartTime.Format(time.StampMilli)
	result += "\n\n"

	result += "### emptyCellIndex ###\n"
	result += fmt.Sprintf("[%d]\n\n", m.emptyCellIndex)

	result += "### cells ###\n"
	result += cellArrayToHtml(&m.cells)
	result += "\n\n"

    result += "### lastTrans ###\n"
	result += m.lastTrans.Markdown()

	return result
}

func (m *Model) String() string {
	result := "state:\n"
	result += m.state.String()

	result += fmt.Sprintf("start time [%s]\n", m.StartTime.Format(time.StampMilli))

	result += "cells ["

	for i := 0; i < NUM_GAME_CELLS; i++ {
		if i > 0 {
			result += ", "
		}
		result += m.cells[i].String()
	}
	result += "]\n"

	result += fmt.Sprintf("empty cell index [%d]\n", m.emptyCellIndex)

	result += "lastTrans:\n"
	result += m.lastTrans.String()
	result += "\n"

	return result
}
