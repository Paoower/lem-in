package objects

import "sort"

// Solution datatype
type Solution struct {
	Paths        []*Path
	PathsTrigger []int
	TotalRooms   int
}

// Constructor
func NewSolution() *Solution {
	return &Solution{
		Paths:        []*Path{},
		PathsTrigger: []int{},
		TotalRooms:   0,
	}
}

func (s *Solution) Sort() {
	sort.Slice(s.Paths, func(i int, j int) bool {
		return len(s.Paths[i].Rooms) < len(s.Paths[j].Rooms)
	})
}

func (s *Solution) GetTriggers() {
	for i := 1; i < len(s.Paths); i++ {
		var trigger int = (i * (len(s.Paths[i].Rooms) - len(s.Paths[i-1].Rooms))) + 1
		if i == 1 {
			trigger++
		} else {
			trigger += s.PathsTrigger[i-2]
		}
		s.PathsTrigger = append(s.PathsTrigger, trigger)
	}
}

func (s *Solution) SetTotalRooms() {
	nbr := 0
	for i := range s.Paths {
		nbr += s.Paths[i].Cost
	}
	s.TotalRooms = nbr
}
