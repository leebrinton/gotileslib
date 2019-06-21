package tileslib

// DEFAULT_SCRAMBLE_ITERATIONS is the default number of random moves that will
// be attempted when scrambling a new game.
const DEFAULT_SCRAMBLE_ITERATIONS = 2500

// CommandModeAType is an enum of the possible command modes.  Values are
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
