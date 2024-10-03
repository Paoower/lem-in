package path

import "lem-in/src/room"

type Path struct {
	Route     []*room.Room
	Available int
}

func NewPath(route []*room.Room, available int) *Path {
	return &Path{
		Route:     route,
		Available: available,
	}
}
