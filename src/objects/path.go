package objects

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
