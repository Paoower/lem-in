package farm

import (
	"lem-in/src/path"
	"lem-in/src/room"
)

// Farm datatype
type Farm struct {
	Ants     int
	PathsCap int
	Rooms    []*room.Room
	Paths    []*path.Path
}

// Constructor
func NewFarm() *Farm {
	return &Farm{
		Ants:  0,
		Rooms: []*room.Room{},
	}
}
