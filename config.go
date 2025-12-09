// Package tileslib core types for tiles puzzle games.
//
// Copyright (C) 2021 H. Lee Brinton.
// License GPLv3+: GNU GPL version 3 or later
// <http://gnu.org/licenses/gpl.html>
// This is free software: you are free to change and redistribute it.
// There is NO WARRANTY, to the extent permitted by law.
//
package tileslib

// DefaultScrambleIterations is the default number of random moves that will
// be attempted when scrambling a new game.
const DefaultScrambleIterations = 2500

// CommandModeType is an enum of the possible command modes.  Values are
// EmptyCellCentric and ValueCellCentric.
type CommandModeType int

const (
	// EmptyCellCentric movement commands move the empty cell
	EmptyCellCentric CommandModeType = iota
	//ValueCellCentric movement commands move a value cell
	ValueCellCentric
)

// Config is a structure that holds configuration option values. The members
// are ScrambleIterations and CommandMode.
// ScrambleIterations is the number of random moves that will be attempted
// when scrambling a new game.
// CommandMode is a CommandModeType that determines how to interpret key
// input.
type Config struct {
	ScrambleIterations int
	CommandMode        CommandModeType
}
