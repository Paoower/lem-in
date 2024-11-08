package farm

import (
	"lem-in/src/path"
	"lem-in/src/solution"
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
		tempName := f.Paths[i].Rooms[1].Name
		if !slices.Contains(sliceOfStartingRoomsName, tempName) {
			sliceOfStartingRoomsName = append(sliceOfStartingRoomsName, tempName)
		}
		tempName = f.Paths[i].Rooms[len(f.Paths[i].Rooms)-2].Name
		if !slices.Contains(sliceOfEndingRoomsName, tempName) {
			sliceOfEndingRoomsName = append(sliceOfEndingRoomsName, tempName)
		}
	}

	start := len(sliceOfStartingRoomsName)
	end := len(sliceOfEndingRoomsName)
	if end > start {
		f.PathsCap = end
	}
	f.PathsCap = start
}

func IsACompatiblePath(solutionSlice solution.Solution, path *path.Path) bool {
	for indexPathInSolution := 0; indexPathInSolution < len(solutionSlice.Paths); indexPathInSolution++ {
		for indexRoomInSolutionPath := 0; indexRoomInSolutionPath < len(solutionSlice.Paths[indexPathInSolution].Rooms); indexRoomInSolutionPath++ {
			for indexRoomInArgPath := 0; indexRoomInArgPath < len(path.Rooms); indexRoomInArgPath++ {
				if path.Rooms[indexRoomInArgPath] == solutionSlice.Paths[indexPathInSolution].Rooms[indexRoomInSolutionPath] {
					return false
				}
			}
		}
	}
	return true
}

func (f *Farm) LookingForEveryPossibleSolution() {
	for indexPath := range f.Paths {
		solutionSlice := f.InializationSolutionSlice(indexPath)
		for nbrOfPaths := 1; nbrOfPaths < f.PathsCap; nbrOfPaths++ {
			for indexSolution := range solutionSlice {
				Sœ
				if len(solutionSlice[indexSolution].Paths) == nbrOfPaths {
					for otherPath := range f.Paths {
						if otherPath != indexPath {
							if IsACompatiblePath(solutionSlice[indexSolution], f.Paths[otherPath]) {
								solutionSlice = append(solutionSlice, solutionSlice[indexSolution])
								solutionSlice[indexSolution].Paths = append(solutionSlice[indexSolution].Paths, f.Paths[otherPath])
							}
						}
					}
				}
			}
		}
		f.Solutions = append(f.Solutions, solutionSlice)
	}
}

// a bloc function that create a slice Of the struct Solution, and initalize it's first path to the one chosen in parameter
func (f *Farm) InializationSolutionSlice(index int) []solution.Solution {
	var solutionSlice []solution.Solution
	firstPath := f.Paths[index]
	firstSolution := solution.NewSolution()
	firstSolution.Paths[0] = firstPath
	solutionSlice = append(solutionSlice, firstSolution)
	return solutionSlice
}

func (f *Farm) SortPaths() {
	f.sortPathSize()
	f.GetPathCap()
	//f.LookingForEveryPossibleSolution()
}
