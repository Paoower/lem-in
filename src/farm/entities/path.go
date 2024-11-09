package entities

import "fmt"

type Path struct {
	Rooms []*Room
	Cost  int
}

func NewPath(rooms []*Room) *Path {
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
