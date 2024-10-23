package farm

import (
	"fmt"
	"slices"
	"sort"
)

func (f *Farm) sortPathSize() {
	f.GetPathCap()
	sort.Slice(f.Paths, func(i int, j int) bool {
		return len(f.Paths[i].Rooms) < len(f.Paths[j].Rooms)
	})
}

func (f *Farm) GetPathCap() {
	var sliceOfStartingRoomsName []string
	var sliceOfEndingRoomsName []string

	for i := 0; i < len(f.Paths); i++ {
		tempName := f.Paths[i].Rooms[0].Name
		if !slices.Contains(sliceOfStartingRoomsName, tempName) {
			sliceOfStartingRoomsName = append(sliceOfStartingRoomsName, tempName)
		}
		tempName = f.Paths[i].Rooms[len(f.Paths[i].Rooms)-1].Name
		if !slices.Contains(sliceOfEndingRoomsName, tempName) {
			sliceOfEndingRoomsName = append(sliceOfEndingRoomsName, tempName)
		}
	}

	fmt.Println(sliceOfStartingRoomsName)
	fmt.Println(sliceOfEndingRoomsName)

	start := len(sliceOfStartingRoomsName)
	end := len(sliceOfEndingRoomsName)
	if end > start {
		f.PathsCap = end
	}
	f.PathsCap = start
}

func (f *Farm) SortPaths() {
	f.sortPathSize()
	f.GetPathCap()
}
