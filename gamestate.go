// Package tileslib core types for tiles puzzle games.
//
// Copyright (C) 2021 H. Lee Brinton.
// License GPLv3+: GNU GPL version 3 or later
// <http://gnu.org/licenses/gpl.html>
// This is free software: you are free to change and redistribute it.
// There is NO WARRANTY, to the extent permitted by law.
//
package tileslib

import "fmt"

// GameSolved is used to determine if a game is solved by comparing it to a GameState.SolvedState.
const GameSolved = uint16(0)

// NumGameCells is the number of cells in the game matrix.
const NumGameCells = 16

// UInt16Max is the maximum value that be held in an unsigned 16 bit integer.
const UInt16Max = uint16(65535)

var _bitOnMasks [NumGameCells]uint16
var _bitOffMasks [NumGameCells]uint16

// GameState is a data structure to hold game level data.
// SolvedState is used to track whether the  game is solved or not.
type GameState struct {
	SolvedState uint16
}

func init() {
	loadBitMasks(&_bitOnMasks, &_bitOffMasks)
}

func twoToTheX(x byte) uint16 {
	switch x {
	case 0:
		return uint16(1)

	case 1:
		return uint16(2)

	default:
		result := 2
		for counter := (x - 1); counter != 0; counter-- {
			result *= 2
		}
		return uint16(result)
	}
	return uint16(0)
}

func loadBitOnMasks(onmasks *[NumGameCells]uint16) {
	for i := 0; i < NumGameCells; i++ {
		onmasks[i] = twoToTheX(byte(i))
	}
}

func loadBitOffMasks(onmasks *[NumGameCells]uint16, offmasks *[NumGameCells]uint16) {
	for i := 0; i < NumGameCells; i++ {
		offmasks[i] = uint16(UInt16Max - onmasks[i])
	}
}

func loadBitMasks(onmasks *[NumGameCells]uint16, offmasks *[NumGameCells]uint16) {
	loadBitOnMasks(onmasks)
	loadBitOffMasks(onmasks, offmasks)
}

// NewGameState creates a new GameState data structure.
func NewGameState() *GameState {
	gs := new(GameState)
	gs.SolvedState = GameSolved
	return gs
}

// SetBitOn set the indexth bit of the state variable on.

func (gs *GameState) SetBitOn(index int) {
	gs.SolvedState = uint16(gs.SolvedState | _bitOnMasks[index])
}

// SetBitOff set the indexth bit of the state variable off.
func (gs *GameState) SetBitOff(index int) {
	gs.SolvedState = uint16(gs.SolvedState & _bitOffMasks[index])
}

// Solved determine if the game is solved. Return TRUE if the player has won
// otherwise return FALSE.
func (gs *GameState) Solved() bool {
	result := false

	if gs.SolvedState == GameSolved {
		result = true
	}
	return result
}

// Update Update the state variable. Given two indexes a cell index and
// a cell value index, If the cell index and the cell value index are equal
// (the cell holds the proper value), then set the cell indexth bit of the
// state variable on otherwise set it off.
func (gs *GameState) Update(currentIndex int, correctIndex int) {
	if currentIndex == correctIndex {
		gs.SetBitOff(currentIndex)
	} else {
		gs.SetBitOn(currentIndex)
	}
}

// SolvedValue returns a value from 0 to 16 that is the number of cells that
// are in their solved position.
func (gs *GameState) SolvedValue() int {
	test := uint16(0)
	result := 0

	for i := 0; i < NumGameCells; i++ {
		test = (gs.SolvedState & _bitOnMasks[i])

		if test == 0 {
			result++
		}
	}
	return result
}

// SolvedPercent returns the percentage of cells that are in their solved
// position.
func (gs *GameState) SolvedPercent() int {
	value := gs.SolvedValue()
	tmp := float32(float32(value) / float32(NumGameCells))

	return int(tmp * 100)
}

// String returns a string representation of a GameState.
func (gs GameState) String() string {
	solved := "No"
	if gs.Solved() {
		solved = "Yes"
	}

	result := "solved ["
	result += solved
	result += "]\n"

	result += "on masks "
	result += fmt.Sprint(_bitOnMasks)
	result += "\n"

	result += "off masks "
	result += fmt.Sprint(_bitOffMasks)
	result += "\n"

	return result
}

func maskArrayToHtml(array *[NumGameCells]uint16) string {
	result := "<table><tr>"

	for i := 0; i < NumGameCells; i++ {
		result += "<td>"
		result += fmt.Sprintf("%d", array[i])
		result += "</td>"
	}
	result += "</tr></table>"
	return result
}

// Markdown	returns a markdown representation of a GameState.
func (gs GameState) Markdown() string {
	solved := "No"
	if gs.Solved() {
		solved = "Yes"
	}

	md := fmt.Sprintf("* Solved? [%s]\n", solved)

	md += "* On Masks "
	md += maskArrayToHtml(&_bitOnMasks)
	md += "\n"

	md += "* Off Masks"
	md += maskArrayToHtml(&_bitOffMasks)
	md += "\n"

	return md
}
