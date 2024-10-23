package farm

import (
	"lem-in/src/path"
	"lem-in/src/room"
)

// Farm datatype
type Farm struct {
	Ants       int
	Rooms      []*room.Room
	PathObject *path.PathObject
}

// Constructor
func NewFarm() *Farm {
	return &Farm{
		Ants:       0,
		Rooms:      []*room.Room{},
		PathObject: &path.PathObject{},
	}
}
