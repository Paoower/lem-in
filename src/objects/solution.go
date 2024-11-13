package objects

import "sort"

// Solution datatype
type Solution struct {
	Paths        []*Path
	PathsTrigger []int
}

// Constructor
func NewSolution() Solution {
	return Solution{
		Paths:        []*Path{},
		PathsTrigger: []int{},
	}
}

func (s *Solution) Sort() {
	sort.Slice(s.Paths, func(i int, j int) bool {
		return len(s.Paths[i].Rooms) < len(s.Paths[j].Rooms)
	})
}

func (s *Solution) GetTriggers() {
	var previousPathsCosts int = len(s.Paths[0].Rooms)
	for i := 1; i < len(s.Paths); i++ {
		var trigger int = (len(s.Paths[i].Rooms)+1)*i + 2 - previousPathsCosts
		s.PathsTrigger = append(s.PathsTrigger, trigger)
	}
}
