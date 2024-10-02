package farm

import (
	"lem-in/src/room"
)

// Farm datatype
type Farm struct {
	Ants  int
	Rooms []*room.Room
}

// Constructor
func NewFarm() *Farm {
	return &Farm{
		Ants:  0,
		Rooms: []*room.Room{},
	}
}
