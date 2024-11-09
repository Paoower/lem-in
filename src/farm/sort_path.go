package farm

import (
	"fmt"
	e "lem-in/src/farm/entities"
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

func IsACompatiblePath(solutionSlice e.Solution, path *e.Path) bool {
	for i := 0; i < len(solutionSlice.Paths); i++ {
		solutionPathRooms := solutionSlice.Paths[i].Rooms
		for j := 0; j < len(solutionPathRooms) - 1; j++ {
			for k := 0; k < len(path.Rooms) - 1; k++ {
				if k != 0 && j != 0 && path.Rooms[k] == solutionPathRooms[j] {
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
func (f *Farm) InializationSolutionSlice(index int) []e.Solution {
	var solutionSlice []e.Solution
	firstPath := f.Paths[index]
	firstSolution := e.NewSolution()
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
