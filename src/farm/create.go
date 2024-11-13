package farm

import (
	"bufio"
	"errors"
	"lem-in/src/objects"
	t "lem-in/src/tools"
	"log"
	"os"
	"strconv"
	"strings"
)

// insert room element into rooms slice
func (farm *Farm) insert(index int, value *objects.Room) {
	if len(farm.Rooms) == index {
		farm.Rooms = append(farm.Rooms, value)
		return
	}
	farm.Rooms = append(farm.Rooms[:index+1], farm.Rooms[index:]...)
	farm.Rooms[index] = value
}

// Main function for the farm's mapping
func (farm *Farm) Create(filepath string) {
	var err error

	// Slice variables
	var links []string

	// Variables to track the first line
	var firstline bool = true

	// Handle file opening
	var file *os.File
	file, err = os.Open(filepath)
	t.Check(err)
	defer file.Close()

	// Variables to get the start and end rooms
	var start bool = false
	var end bool = false
	var finalRoom *objects.Room

	// Scanner to read the file
	var scanner *bufio.Scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		var line string = scanner.Text()

		// Skip comments
		if strings.HasPrefix(line, "#") && !(strings.Contains(line, "start") || strings.Contains(line, "end")) {
			continue
		}

		// Get ants number on the first line
		if firstline {
			var e error
			var a int

			a, e = strconv.Atoi(line)
			if e != nil {
				log.Fatal(errors.New("invalid number of ants"))
			}

			// Check ants count
			t.Check(farm.checkAnts(a))

			// Set ants value and skip line
			farm.TotalAnts = a
			firstline = false
			continue
		}

		// Check for start room
		if strings.HasPrefix(line, "##sta") {
			t.Debug("Found start room")
			start = true
			continue
		}

		// Check for end room
		if strings.HasPrefix(line, "##en") {
			end = true
			continue
		}

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Add links
		if strings.Contains(line, "-") {
			if farm.checkRoomName(line) == nil {
				links = append(links, line)
				continue
			} else {
				break
			}
		}

		// Parse room data
		var room objects.Room
		var e error
		room, e = farm.parseRoom(line)
		t.Check(e)

		// Add the room at the beginning of the slice if its the start room
		if start {
			farm.insert(0, &room)
			start = false
			continue
		}

		// Add the room at the end of the slice if its the final room
		if end {
			finalRoom = &room
			end = false
			continue
		}

		// Adding parsed room data to the slice
		farm.Rooms = append(farm.Rooms, &room)
	}

	// Add final room
	if finalRoom != nil {
		farm.Rooms = append(farm.Rooms, finalRoom)
	}

	// Parsing links
	err = farm.parseLinks(links)
	t.Check(err)

	// Check if the first/final room is reachable
	t.Check(farm.Rooms[0].CheckRoomLink())
	t.Check(farm.Rooms[len(farm.Rooms)-1].CheckRoomLink())

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
