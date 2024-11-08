package farm

import (
	o "lem-in/src/objects"
)

// Farm datatype
type Farm struct {
	TotalAnts	int
	Ants		[]*o.Ant
	AntNb		int
	PathsCap	int
	Rooms		[]*o.Room
	Paths		[]*o.Path
	Solutions	[]*o.Solution

}

// Constructor
func NewFarm() *Farm {
	return &Farm{
		TotalAnts:	0,
		Ants:		[]*o.Ant{},
		AntNb:		0,
		Rooms:		[]*o.Room{},
	}
}
