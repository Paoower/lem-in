package entities

import (
	"fmt"
	"testing"
)

func getTestPath() *Path {
	rooms := []*Room{
		NewRoom("a", 0, 0),
		NewRoom("b", 0, 0),
		NewRoom("end", 0, 0),
	}
	return NewPath(rooms)
}

func TestBasicMoveAnt(t *testing.T) {
	path := getTestPath()
	rooms := path.Rooms
	ant := NewAnt(1, path)
	rooms[0].Ants = append(rooms[0].Ants, ant)
	status := ant.Move()
	fmt.Printf("status: %s", status.Message())
	if len(rooms[0].Ants) > 0 {
		t.Fatalf("The ant is still in the room 1.")
	}
	if len(rooms[1].Ants) == 0 {
		t.Fatalf("The ant failed to move in room 2.")
	}
	if ant.IndexRoom != 1 {
		t.Fatalf("The ant has the wrong index room.")
	}
	ant.Move() // to end
	ant.Move() // to finish (deleted)
	for i, room := range rooms {
		if len(room.Ants) != 0 {
			t.Fatalf("The ant is in the room of index %d and should not be.", i)
		}
	}
}

func TestCollisionMoveAnt(t *testing.T) {
	path := getTestPath()
	rooms := path.Rooms
	ant1 := NewAnt(1, path)
	ant2 := NewAnt(2, path)
	rooms[0].Ants = append(rooms[0].Ants, ant1)
	ant1.Move() // move into b
	rooms[0].Ants = append(rooms[0].Ants, ant2)
	ant2.Move() // try to move into b
	if len(rooms[0].Ants) == 0 || rooms[0].Ants[0].Id != 2 {
		t.Fatalf("The ant 2 is not in the first room")
	}
	if len(rooms[1].Ants) > 1 {
		t.Fatalf("There are more than 1 ants in the room.")
	}
	if rooms[1].Ants[0].Id != 1 {
		t.Fatalf("The ant in the room is not the right one.")
	}
	ant1.Move() // move into end
	ant2.Move() // move into b
	ant2.Move() // move into end
	endAntsNb := len(rooms[2].Ants)
	if endAntsNb != 2 {
		t.Fatalf("There are %d ant in end room, but it should be 2.", endAntsNb)
	}
}
