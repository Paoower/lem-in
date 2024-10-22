package farm

import (
	"lem-in/src/path"
	"lem-in/src/room"
	"slices"
)

// Farm datatype
type Farm struct {
	Ants  int
	Rooms []*room.Room
	Paths []*path.Path
}

// Constructor
func NewFarm() *Farm {
	return &Farm{
		Ants:  0,
		Rooms: []*room.Room{},
		Paths: []*path.Path{},
	}
}

func (f *Farm) GetPathCap() int {
	var sliceOfStartingRoomsName []string
	var sliceOfEndingRoomsName []string

	for i := 0; i < len(f.Paths); i++ {
		tempName := f.Paths[i].Route[0].Name
		if !slices.Contains(sliceOfStartingRoomsName, tempName) {
			sliceOfStartingRoomsName = append(sliceOfStartingRoomsName, tempName)
		}
		tempName = f.Paths[i].Route[len(f.Paths[i].Route)-1].Name
		if !slices.Contains(sliceOfEndingRoomsName, tempName) {
			sliceOfEndingRoomsName = append(sliceOfEndingRoomsName, tempName)
		}
	}

	start := len(sliceOfStartingRoomsName)
	end := len(sliceOfEndingRoomsName)
	if end > start {
		return end
	}
	return start
}
