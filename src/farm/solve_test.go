package farm

import (
	e "lem-in/src/farm/entities"
	"testing"
)

func createTestFarm() (*Farm, *e.Solution) {
	farm := NewFarm()
	solution := e.NewSolution()
	rooms := []*e.Room{
		e.NewRoom("a", 0, 0),
		e.NewRoom("b", 0, 0),
		e.NewRoom("end", 0, 0),
	}
	farm.Rooms = rooms
	farm.TotalAnts = 10
	path1 := e.NewPath(append([]*e.Room{}, rooms[0], rooms[2]))
	path2 := e.NewPath(append([]*e.Room{}, rooms[1], rooms[2]))
	solution.Paths = append(solution.Paths, path1, path2)
	return farm, &solution
}

func TestAddNewAnts(t *testing.T) {
	farm, solution := createTestFarm()
	antsAdded := farm.addNewAnts(solution)
	if (antsAdded != 2) {
		t.Fatalf("%d ants were added, not 2.", antsAdded)
	}
	if (farm.AntNb != 2) {
		t.Fatalf("farm.AntNb=%d, not 2.", farm.AntNb)
	}
	if (len(farm.Ants) != 2) {
		t.Fatalf("%d ants are stored in farm.Ants, not 2.", len(farm.Ants))
	}
	endRoom := farm.Rooms[len(farm.Rooms) - 1]
	if (len(endRoom.Ants) != 0) {
		t.Fatalf("ant %d should not be in end room.", endRoom.Ants[0].Id)
	}
	antsAdded = farm.addNewAnts(solution)
	if (antsAdded != 0) {
		t.Fatalf("%d ants were added into full rooms.", antsAdded)
	}
}

func TestMoveAnts(t *testing.T) {
	farm, solution := createTestFarm()
	farm.addNewAnts(solution)
	farm.moveCurrentsAnts()
	endRoom := farm.Rooms[len(farm.Rooms) - 1]
	if len(endRoom.Ants) != 2 {
		t.Fatalf("%d are in end room, not 2.", len(endRoom.Ants))
	}
	farm.moveCurrentsAnts()
	if len(endRoom.Ants) != 0 {
		t.Fatalf("ant %d should not be in end room.", endRoom.Ants[0].Id)
	}
}
