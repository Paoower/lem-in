package solution

import "lem-in/src/path"

// Solution datatype
type Solution struct {
	Paths        []*path.Path
	PathsTrigger []int
}

// Constructor
func NewSolution() Solution {
	return Solution{
		Paths:        []*path.Path{},
		PathsTrigger: []int{},
	}
}
