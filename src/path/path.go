package path

import "lem-in/src/room"

type PathObject struct {
	PathsCap     int
	Paths        []*Path
	PathTriggers []int
}

type Path struct {
	Rooms []*room.Room
}

func NewPathObject() *PathObject {
	return &PathObject{
		PathsCap:     0,
		Paths:        []*Path{},
		PathTriggers: []int{},
	}
}

func NewPath(rooms []*room.Room) *Path {
	return &Path{
		Rooms: rooms,
	}
}
