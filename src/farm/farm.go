package farm

import e "lem-in/src/farm/entities"

// Farm datatype
type Farm struct {
	TotalAnts	int
	Ants		[]*e.Ant
	AntNb		int
	PathsCap	int
	Rooms		[]*e.Room
	Paths		[]*e.Path
	Solutions	[]e.Solution

}

// Constructor
func NewFarm() *Farm {
	return &Farm{
		TotalAnts:	0,
		Ants:		[]*e.Ant{},
		AntNb:		0,
		Rooms:		[]*e.Room{},
	}
}
