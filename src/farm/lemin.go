package farm

import (
	"fmt"
	e "lem-in/src/farm/entities"
	"slices"
)

func	(farm *Farm) PrintAntsPositions() {
	var ant			*e.Ant
	var currentRoom	*e.Room
	var i			int
	var antsLen		int

	antsLen = len(farm.Ants)
	for i, ant = range farm.Ants {
		currentRoom = ant.Path.Rooms[ant.IndexRoom]
		fmt.Printf("L%d-%s", ant.Id, currentRoom.Name)
		if i == antsLen - 1 {
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
func	(farm *Farm) AddNewAnts(solution *e.Solution) int {
	var path		*e.Path
	var ant			*e.Ant
	var firstRoom	*e.Room
	var i			int
	var antAdded	int

	antAdded = 0
	for i = 0; i < len(solution.Paths); i++ {
		path = solution.Paths[i]
		if farm.AntNb == farm.TotalAnts {
			return antAdded
		} else if path == nil || len(path.Rooms) == 0 {
			continue
		}
		firstRoom = path.Rooms[0]
		if len(firstRoom.Ants) > 0 && len(path.Rooms) > 1 {
			// room already used and is not the end
			continue
		}
		ant = e.NewAnt(farm.AntNb + 1, path)
		farm.AntNb++
		farm.Ants = append(farm.Ants, ant)
		firstRoom.Ants = append(firstRoom.Ants, ant)
		antAdded++
	}
	return antAdded
}

func	(farm *Farm) MoveCurrentsAnts() {
	var i			int
	var antStatus	e.AntStatus

	i = 0
	for i < len(farm.Ants) {
		antStatus = farm.Ants[i].Move()
		if antStatus == e.AntStatusNotValid {
			panic(fmt.Sprintf("Error: Ant %d not valid", farm.Ants[i].Id))
		}
		if antStatus == e.AntStatusDeleted {
			farm.Ants = slices.Delete(farm.Ants, i, i+1)
			continue
		}
		i++
	}
}

func	(farm *Farm) LemIn() {
	var solution	e.Solution

	for {
		// find the solution
		farm.MoveCurrentsAnts()
		farm.AddNewAnts(&solution)
		if len(farm.Ants) == 0 {
			break
		}
		farm.PrintAntsPositions()
	}
}
