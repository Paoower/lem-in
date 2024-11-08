package farm

import (
	"lem-in/src/path"
	"lem-in/src/room"
	"lem-in/src/solution"
)

// Farm datatype
type Farm struct {
	Ants      int
	PathsCap  int
	Rooms     []*room.Room
	Paths     []*path.Path
	Solutions []solution.Solution
}

// Constructor
func NewFarm() *Farm {
	return &Farm{
		Ants:  0,
		Rooms: []*room.Room{},
	}
}
