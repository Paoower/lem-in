package farm

import (
	"fmt"
	"slices"
	"sort"
)

func (f *Farm) SortPaths() {
	f.GetPathCap()
	sort.Slice(f.PathObject.Paths, func(i int, j int) bool {
		return len(f.PathObject.Paths[i].Rooms) < len(f.PathObject.Paths[j].Rooms)
	})
}

func (f *Farm) GetPathCap() {
	var sliceOfStartingRoomsName []string
	var sliceOfEndingRoomsName []string

	for i := 0; i < len(f.PathObject.Paths); i++ {
		tempName := f.PathObject.Paths[i].Rooms[1].Name
		if !slices.Contains(sliceOfStartingRoomsName, tempName) {
			sliceOfStartingRoomsName = append(sliceOfStartingRoomsName, tempName)
		}
		tempName = f.PathObject.Paths[i].Rooms[len(f.PathObject.Paths[i].Rooms)-2].Name
		if !slices.Contains(sliceOfEndingRoomsName, tempName) {
			sliceOfEndingRoomsName = append(sliceOfEndingRoomsName, tempName)
		}
	}

	fmt.Println(sliceOfStartingRoomsName)
	fmt.Println(sliceOfEndingRoomsName)

	start := len(sliceOfStartingRoomsName)
	end := len(sliceOfEndingRoomsName)
	if end > start {
		f.PathObject.PathsCap = end
	}
	f.PathObject.PathsCap = start
}
