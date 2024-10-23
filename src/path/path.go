package path

import "lem-in/src/room"

type Path struct {
	Rooms []*room.Room
	Cost  int
}

func NewPath(rooms []*room.Room) *Path {
	return &Path{
		Rooms: rooms,
		Cost:  len(rooms),
	}
}

// Solution datatype
type Solution struct {
	Paths        []*Path
	PathsTrigger []int
}

// Constructor
func NewSolution() *Solution {
	return &Solution{
		Paths:        []*Path{},
		PathsTrigger: []int{},
	}
}
