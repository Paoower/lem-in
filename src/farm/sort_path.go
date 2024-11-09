package farm

import (
	"fmt"
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
		for indexRoomInSolutionPath := 0; indexRoomInSolutionPath < len(solutionSlice.Paths[indexPathInSolution].Rooms)-1; indexRoomInSolutionPath++ {
			for indexRoomInArgPath := 0; indexRoomInArgPath < len(path.Rooms)-1; indexRoomInArgPath++ {
				if indexRoomInArgPath != 0 && indexRoomInSolutionPath != 0 && path.Rooms[indexRoomInArgPath] == solutionSlice.Paths[indexPathInSolution].Rooms[indexRoomInSolutionPath] {
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
		f.Solutions = append(f.Solutions, solutionSlice...)
	}
}

// a bloc function that create a slice Of the struct Solution, and initalize it's first path to the one chosen in parameter
func (f *Farm) InializationSolutionSlice(index int) []solution.Solution {
	var solutionSlice []solution.Solution
	firstPath := f.Paths[index]
	firstSolution := solution.NewSolution()
	firstSolution.Paths = append(firstSolution.Paths, firstPath)
	solutionSlice = append(solutionSlice, firstSolution)
	return solutionSlice
}

func (f *Farm) TestCheckingForAllSolutions() {
	for _, s := range f.Solutions {
		fmt.Println(s)
	}
}

func (f *Farm) GetRidOfCopy() {
	for indexFirstSolution := range f.Solutions {
		for indexSecondSolution := indexFirstSolution + 1; indexSecondSolution < len(f.Solutions); indexSecondSolution++ {
			if len(f.Solutions[indexFirstSolution].Paths) == len(f.Solutions[indexSecondSolution].Paths) {
				isaMatch := false
				for indexPathInFirstSolution := range f.Solutions[indexFirstSolution].Paths {
					if isaMatch || indexPathInFirstSolution == 0 {
						isaMatch = false
						for indexPathInSecondSolution := range f.Solutions[indexSecondSolution].Paths {
							if f.Solutions[indexFirstSolution].Paths[indexPathInFirstSolution] == f.Solutions[indexSecondSolution].Paths[indexPathInSecondSolution] {
								isaMatch = true
							}
						}
					}
				}
				if isaMatch {
					if indexSecondSolution == len(f.Solutions)-1 {
						f.Solutions = f.Solutions[:indexSecondSolution]
					} else {
						f.Solutions = append(f.Solutions[:indexSecondSolution], f.Solutions[indexSecondSolution+1:]...)
						indexSecondSolution--
					}
				}
			}
		}
	}
}

func (f *Farm) sortsolutions() {
	for _, s := range f.Solutions {
		s.Sort()
		s.GetTriggers()
		fmt.Println(s)
	}
}

func (f *Farm) SortPaths() {
	f.sortPathSize()
	f.GetPathCap()
	f.LookingForEveryPossibleSolution()
	f.GetRidOfCopy()
	f.sortsolutions()
}
