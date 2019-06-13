package tileslib

const DEFAULT_SCRAMBLE_ITERATIONS = 2500

type CommandModeType int

const (
	EmptyCellCentric CommandModeType = iota
	ValueCellCentric
)

type Config struct {
	ScrambleIterations int
	CommandMode CommandModeType	
}

