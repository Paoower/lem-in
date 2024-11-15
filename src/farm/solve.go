package farm

import (
	"fmt"
	o "lem-in/src/objects"
	"slices"
)

func (farm *Farm) printAntsPositions() {
	var ant *o.Ant
	var currentRoom *o.Room
	var i int
	var antsLen int

	antsLen = len(farm.Ants)
	for i, ant = range farm.Ants {
		currentRoom = ant.Path.Rooms[ant.IndexRoom]
		fmt.Printf("L%d-%s", ant.Id, currentRoom.Name)
		if i == antsLen-1 {
			// on last element
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}

// Add awaiting ants into the circuit.
//
// Returns the number of ants added.
func (farm *Farm) addNewAnts(solution *o.Solution) int {
	var path *o.Path
	var ant *o.Ant
	var firstRoom *o.Room
	var i int
	var antAdded int

	antAdded = 0
	for i = 0; i < len(solution.Paths); i++ {
		path = solution.Paths[i]
		if farm.AntNb == farm.TotalAnts {
			return antAdded
		} else if path == nil || len(path.Rooms) <= 1 {
			continue
		}
		firstRoom = path.Rooms[1] // ignore start room
		if len(firstRoom.Ants) > 0 && len(path.Rooms) > 2 {
			// room already used and is not the end
			continue
		}
		ant = o.NewAnt(farm.AntNb+1, path)
		farm.AntNb++
		farm.Ants = append(farm.Ants, ant)
		firstRoom.Ants = append(firstRoom.Ants, ant)
		antAdded++
	}
	return antAdded
}

func (farm *Farm) moveCurrentsAnts() {
	var i int
	var antStatus o.AntStatus

	i = 0
	for i < len(farm.Ants) {
		antStatus = farm.Ants[i].Move()
		if antStatus == o.AntStatusNotValid {
			panic(fmt.Sprintf("Error: Ant %d not valid", farm.Ants[i].Id))
		}
		if antStatus == o.AntStatusDeleted {
			farm.Ants = slices.Delete(farm.Ants, i, i+1)
			continue
		}
		i++
	}
}

func (farm *Farm) selectSolution() *o.Solution {
	for i := len(farm.Solutions) - 1; i > 0; i-- {
		if farm.Solutions[i].PathsTrigger[len(farm.Solutions[i].PathsTrigger)-1] <= (farm.TotalAnts - farm.AntNb) {
			return farm.Solutions[i]
		}
	}
	return farm.Solutions[0]
}

func (farm *Farm) Solve() {
	var solution *o.Solution
	cpt := 0
	for {
		solution = farm.selectSolution()
		farm.moveCurrentsAnts()
		farm.addNewAnts(solution)
		if len(farm.Ants) == 0 {
			break
		}
		cpt++
		farm.printAntsPositions()
	}
	fmt.Println(cpt)
}
