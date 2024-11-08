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

	for i, path = range solution.Paths {
		if farm.AntNb == farm.TotalAnts {
			return i
		}
		if path == nil {
			continue
		}
		if len(path.Rooms) < 1 {
			continue
		}
		firstRoom = path.Rooms[0]
		if len(firstRoom.Ants) > 0 && len(path.Rooms) > 1 {
			// room already used and is not the end
			continue
		}
		ant = e.NewAnt(farm.AntNb + 1, path) // create ant
		farm.AntNb += 1
		farm.Ants = append(farm.Ants, ant) // add ant into the slice of ants
		firstRoom.Ants = append(firstRoom.Ants, ant)
		// add ant into the first room
	}
	return i
}

func	(farm *Farm) MoveCurrentsAnts() {
	var i			int
	var antStatus	e.AntStatus

	i = 0
	for i < len(farm.Ants) {
		antStatus = farm.Ants[i].Move()
		if antStatus == e.AntNotValid {
			panic(fmt.Sprintf("Error: Ant %d not valid", farm.Ants[i].Id))
		}
		if antStatus == e.AntDeleted {
			farm.Ants = append(farm.Ants[:i], farm.Ants[i+1:]...)
			farm.Ants = slices.Delete(farm.Ants, i, i+1)
			continue
		}
		i++
	}
}
