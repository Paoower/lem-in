package path

import "lem-in/src/room"

type PathObject struct {
	PathsCap int
	Paths    []*Path
}

type Path struct {
	Rooms []*room.Room
	Cost  int
}

func NewPathObject() *PathObject {
	return &PathObject{
		PathsCap: 0,
		Paths:    []*Path{},
	}
}

func NewPath(rooms []*room.Room) *Path {
	return &Path{
		Rooms: rooms,
		Cost:  len(rooms),
	}
}
