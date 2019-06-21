package tileslib

import (
	"math/rand"
	"time"
)

var _random *rand.Rand

// Direction is an enum for the four directions Up, Doen, Left and Right.
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func init() {
	now := time.Now()
	source := rand.NewSource(now.UnixNano())
	_random = rand.New(source)
}

// String return a string representation of a Direction.
func (d Direction) String() string {
	result := ""

	switch d {
	case Up:
		result = "Up"

	case Down:
		result = "Down"

	case Left:
		result = "Left"

	case Right:
		result = "Right"
	}
	return result
}

// RandomDirection returns a randomly generated Direction.
func RandomDirection() Direction {
	raw := _random.Intn(4)
	d := Up

	switch raw {
	case 0:
		d = Up

	case 1:
		d = Down

	case 2:
		d = Left

	case 3:
		d = Right
	}
	return d
}

// ReverseDirection returns the opposite Direction.
func ReverseDirection(d Direction) Direction {
	result := d

	switch d {
	case Right:
		result = Left

	case Left:
		result = Right

	case Down:
		result = Up

	case Up:
		result = Down
	}
	return result
}
