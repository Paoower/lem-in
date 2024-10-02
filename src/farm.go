package src

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Room datatype
type Room struct {
	name  string
	x     int
	y     int
	links []*Room
}

// Error helping function
func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Check if room was already saved
func isRoomThere(room Room, rooms []*Room) bool {
	for _, r := range rooms {
		if r.name == room.name {
			return true
		}
	}
	return false
}

// Given a line, this function extracts the room's values
func parseRoom(line string, rooms []*Room) (Room, error) {
	fields := strings.Split(line, " ")

	// Checking the line's format
	if len(fields) != 3 {
		return Room{}, errors.New("ERROR: Incorrect room format")
	}

	// Room's name
	name := fields[0]

	// Converting the first field
	x, err := strconv.Atoi(fields[1])
	if err != nil {
		return Room{}, errors.New("ERROR: Invalid data format for x-coordinate")
	}

	// Converting the second field
	y, err := strconv.Atoi(fields[2])
	if err != nil {
		return Room{}, errors.New("ERROR: Invalid data format for y-coordinate")
	}

	// Creating the room
	r := Room{name, x, y, []*Room{}}

	// Checking if the room is already saved
	if isRoomThere(r, rooms) {
		return Room{}, errors.New("ERROR: Duplicate room")
	}

	return r, nil
}

// Fetches a room based on its name
func getRoom(name string, rooms []*Room) (*Room, error) {
	for _, room := range rooms {
		if room.name == name {
			return room, nil
		}
	}
	return nil, errors.New("ERROR: Invalid data format, room not found")
}

// Establishes the links between different rooms
func parseLinks(rooms []*Room, links []string) error {
	for _, link := range links {
		roomNames := strings.Split(link, "-")

		// Get the first room
		r1, e := getRoom(roomNames[0], rooms)
		if e != nil {
			return errors.New("ERROR: Room doesn't exist")
		}

		// Get the second room
		r2, e := getRoom(roomNames[1], rooms)
		if e != nil {
			return errors.New("ERROR: Room doesn't exist")
		}

		// Case when the room is linked to itself
		if r1 == r2 {
			return errors.New("ERROR: Room is linked to itself")
		}

		// Connect the two rooms
		r1.links = append(r1.links, r2)
		r2.links = append(r2.links, r1)
	}

	return nil
}

// Display all the rooms and their links based on a given slice
func DisplayRooms(rooms []*Room) {
	for _, room := range rooms {
		fmt.Printf("Room: %s \n Links: ", room.name)
		for _, link := range room.links {
			fmt.Printf("%s ", link.name)
		}
		fmt.Println()
	}
}

// Main function for the mapping of the farm
func FarmMapping(filepath string) []*Room {
	// Log flags
	log.SetFlags(0)

	// Slice variables
	rooms := []*Room{}
	links := []string{}

	// Handle file opening
	file, err := os.Open(filepath)
	checkError(err)
	defer file.Close()

	// Scanner to read the file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Skip comment, start or end lines
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Add links
		if strings.Contains(line, "-") {
			links = append(links, line)
			continue
		}

		// Parse room data
		room, err := parseRoom(line, rooms)
		checkError(err)

		// Adding parsed room data to the slice
		rooms = append(rooms, &room)
	}

	// Parsing links
	e := parseLinks(rooms, links)
	checkError(e)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rooms
}
