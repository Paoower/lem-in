package path

import (
	"fmt"
	"lem-in/src/room"
)

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

func (p *Path) Print() {
	for _, r := range p.Rooms {
		fmt.Println("= " + r.Name)
	}
}
