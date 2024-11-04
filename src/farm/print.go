package farm

import (
	"fmt"
	"lem-in/src/room"
)

func	orderByAntNb(rooms []*room.Room) {
	var i		int
	var j		int
	var temp	*room.Room

	i = 1
	for i < len(rooms) {
		j = i
		for j > 0 && rooms[j - 1].AntNb > rooms[j].AntNb {
			temp = rooms[j]
			rooms[j] = rooms[j - 1]
			rooms[j - 1] = temp
			j--
		}
		i++
	}
}

func	(f *Farm) Print() {
	var rooms		[]*room.Room
	var	roomsLen	int

	roomsLen = len(f.Rooms[1:])
	rooms = make([]*room.Room, roomsLen)
	copy(rooms, f.Rooms[1:])
	orderByAntNb(rooms)
	for i, r := range rooms {
		if (r.AntNb == 0) {
			continue
		}
		fmt.Printf("L%d-%s", r.AntNb, r.Name)
		if (i != roomsLen - 1) {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}
