package tileslib

// NewGameAnswer in an enum of values indicating whether or not to start a new
// game.
type NewGameAnswer int

const (
	Accept NewGameAnswer = iota
	Quit
)
