package farm

import "lem-in/src/objects"

// Farm datatype
type Farm struct {
	TotalAnts int
	Ants      []*objects.Ant
	AntNb     int
	PathsCap  int
	Rooms     []*objects.Room
	Paths     []*objects.Path
	Solutions []*objects.Solution
}

// Constructor
func NewFarm() *Farm {
	return &Farm{
		TotalAnts: 0,
		Ants:      []*objects.Ant{},
		AntNb:     0,
		Rooms:     []*objects.Room{},
	}
}
