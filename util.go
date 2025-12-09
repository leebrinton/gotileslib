// Package tileslib core types for tiles puzzle games.
//
// Copyright (C) 2021 H. Lee Brinton.
// License GPLv3+: GNU GPL version 3 or later
// <http://gnu.org/licenses/gpl.html>
// This is free software: you are free to change and redistribute it.
// There is NO WARRANTY, to the extent permitted by law.
//
package tileslib

// NewGameAnswer is an enum of values indicating whether or not to start a new
// game.
type NewGameAnswer int

// Accept prepare a new game
// Quit quit the game
const (
	Accept NewGameAnswer = iota
	Quit
)
