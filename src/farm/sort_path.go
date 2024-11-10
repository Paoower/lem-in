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
	var startingRoomsNames []string
	var endingRoomsNames []string

	for i := 0; i < len(f.Paths); i++ {
		tempName := f.Paths[i].Rooms[1].Name
		if !slices.Contains(startingRoomsNames, tempName) {
			startingRoomsNames = append(startingRoomsNames, tempName)
		}
		tempName = f.Paths[i].Rooms[len(f.Paths[i].Rooms)-2].Name
		if !slices.Contains(endingRoomsNames, tempName) {
			endingRoomsNames = append(endingRoomsNames, tempName)
		}
	}
	start := len(startingRoomsNames)
	end := len(endingRoomsNames)
	if end > start {
		f.PathsCap = end
	}
	f.PathsCap = start
}

func IsACompatiblePath(solutions e.Solution, path *e.Path) bool {
	for i := 0; i < len(solutions.Paths); i++ {
		solutionPathRooms := solutions.Paths[i].Rooms
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
	for i := range f.Paths {
		solutions := f.InializationSolutionSlice(i)
		for nbrOfPaths := 1; nbrOfPaths < f.PathsCap; nbrOfPaths++ {
			for j := range solutions {
				if len(solutions[j].Paths) != nbrOfPaths {
					continue
				}
				for k := range f.Paths {
					if k == i || !IsACompatiblePath(solutions[j], f.Paths[k]) {
						continue
					}
					solutions = append(solutions, solutions[j])
					solutions[j].Paths = append(solutions[j].Paths, f.Paths[k])
				}
			}
		}
		f.Solutions = append(f.Solutions, solutions...)
	}
}

// a bloc function that create a slice Of the struct Solution,
// and initalize it's first path to the one chosen in parameter
func (f *Farm) InializationSolutionSlice(index int) []e.Solution {
	var solutions []e.Solution

	firstPath := f.Paths[index]
	firstSolution := e.NewSolution()
	firstSolution.Paths = append(firstSolution.Paths, firstPath)
	solutions = append(solutions, firstSolution)
	return solutions
}

func (f *Farm) TestCheckingForAllSolutions() {
	for _, s := range f.Solutions {
		fmt.Println(s)
	}
}

func	isAMatch(curSolutionPaths []*e.Path,
					nextSolutionPaths []*e.Path) bool {
	var isAMatch	bool

	isAMatch = false
	for i := range curSolutionPaths {
		if !isAMatch && i != 0 {
			continue
		}
		isAMatch = false
		for j := range nextSolutionPaths {
			if curSolutionPaths[i] == nextSolutionPaths[j] {
				isAMatch = true
			}
		}
	}
	return isAMatch
}

func (f *Farm) GetRidOfCopy() {
	var curSolutionPaths	[]*e.Path
	var nextSolutionPaths	[]*e.Path

	for i := range f.Solutions {
		for j := i + 1; j < len(f.Solutions); j++ {
			curSolutionPaths = f.Solutions[i].Paths
			nextSolutionPaths = f.Solutions[j].Paths
			if len(curSolutionPaths) != len(nextSolutionPaths) {
				continue
			}
			if (!isAMatch(curSolutionPaths, nextSolutionPaths)) {
				continue
			}
			if j == len(f.Solutions)-1 {
				f.Solutions = f.Solutions[:j]
			} else {
				f.Solutions = append(f.Solutions[:j], f.Solutions[j+1:]...)
				j--
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
