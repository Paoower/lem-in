package farm

import (
	"errors"
	"fmt"
	"lem-in/src/room"
	"lem-in/src/tools"
	"strconv"
	"strings"
)

// Given a line, this function extracts the room's values
func (farm *Farm) parseRoom(line string) (room.Room, error) {
	var fields []string = strings.Split(line, " ")
	var e error

	// tools.Checking the line's format
	if len(fields) != 3 {
		fmt.Println(fields)
		return room.Room{}, errors.New("incorrect room format")
	}

	// Room's name
	var name string = fields[0]
	tools.Check(farm.CheckRoomName(name))

	// Converting the first field
	var x int
	x, e = strconv.Atoi(fields[1])
	if e != nil {
		return room.Room{}, errors.New("invalid data format for x-coordinate")
	}

	// Converting the second field
	var y int
	y, e = strconv.Atoi(fields[2])
	if e != nil {
		return room.Room{}, errors.New("invalid data format for y-coordinate")
	}

	// Creating the room
	var r room.Room = room.Room{Name: name, X: x, Y: y, Links: []*room.Room{}}

	// tools.Checking if the room is already saved
	if farm.isRoomThere(r) {
		fmt.Println(r)
		return room.Room{}, errors.New("duplicate room")
	}

	return r, nil
}

// Establishes the links between different rooms
func (farm *Farm) parseLinks(links []string) error {
	for _, link := range links {
		var e error
		roomNames := strings.Split(link, "-")

		// Get the first room
		var r1 *room.Room
		r1, e = farm.getRoom(roomNames[0])
		tools.Check(e)

		// Get the second room
		var r2 *room.Room
		r2, e = farm.getRoom(roomNames[1])
		tools.Check(e)

		// Case when the room is linked to itself
		tools.Check(e)

		// Connect the two rooms
		r1.Links = append(r1.Links, r2)
		r2.Links = append(r2.Links, r1)
	}

	return nil
}
