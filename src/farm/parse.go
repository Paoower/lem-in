package farm

import (
	"errors"
	"fmt"
	e "lem-in/src/farm/entities"
	"lem-in/src/tools"
	"strconv"
	"strings"
)

// Given a line, this function extracts the room's values
func (farm *Farm) parseRoom(line string) (e.Room, error) {
	var fields []string = strings.Split(line, " ")
	var err error

	// tools.Checking the line's format
	if len(fields) != 3 {
		fmt.Println(fields)
		return e.Room{}, errors.New("incorrect room format")
	}

	// Room's name
	var name string = fields[0]
	tools.Check(farm.checkRoomName(name))

	// Converting the first field
	var x int
	x, err = strconv.Atoi(fields[1])
	if err != nil {
		return e.Room{}, errors.New("invalid data format for x-coordinate")
	}

	// Converting the second field
	var y int
	y, err = strconv.Atoi(fields[2])
	if err != nil {
		return e.Room{}, errors.New("invalid data format for y-coordinate")
	}

	// Creating the room
	var r e.Room = e.Room{Name: name, X: x, Y: y, Links: []*e.Room{}}

	// tools.Checking if the room is already saved
	if farm.isRoomThere(r) {
		fmt.Println(r)
		return e.Room{}, errors.New("duplicate room")
	}

	return r, nil
}

// Establishes the links between different rooms
func (farm *Farm) parseLinks(links []string) error {
	for _, link := range links {
		var err error
		roomNames := strings.Split(link, "-")

		// Get the first room
		var r1 *e.Room
		r1, err = farm.getRoom(roomNames[0])
		tools.Check(err)

		// Get the second room
		var r2 *e.Room
		r2, err = farm.getRoom(roomNames[1])
		tools.Check(err)

		// Case when the room is linked to itself
		tools.Check(err)

		// Connect the two rooms
		r1.Links = append(r1.Links, r2)
		r2.Links = append(r2.Links, r1)
	}
	return nil
}
