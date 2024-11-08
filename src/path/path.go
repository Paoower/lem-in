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
