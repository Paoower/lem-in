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

// Type to store a room's information
type Room struct {
	name  string
	x     int
	y     int
	links []*Room
}

// Checks for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Given a line, this function extracts the room's values
func parseRoom(line string) Room {
	fields := strings.Split(line, " ")
	name := fields[0]
	x, _ := strconv.Atoi(fields[1])
	y, _ := strconv.Atoi(fields[2])
	return Room{name, x, y, []*Room{}}
}

// Fetches a room based on its name
func getRoom(name string, rooms []*Room) (*Room, error) {
	for _, room := range rooms {
		if room.name == name {
			return room, nil
		}
	}
	return nil, errors.New("Room not found")
}

// Establishes the links between different rooms
func parseLinks(rooms []*Room, links []string) {
	for _, link := range links {
		roomNames := strings.Split(link, "-")
		r1, e := getRoom(roomNames[0], rooms)
		check(e)
		r2, e := getRoom(roomNames[1], rooms)
		check(e)

		r1.links = append(r1.links, r2)
		r2.links = append(r2.links, r1)
	}
}

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
	rooms := []*Room{}
	links := []string{}
	file, err := os.Open(filepath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "##") {
			continue
		}

		if strings.Contains(line, "-") {
			links = append(links, line)
			continue
		}

		room := parseRoom(line)
		rooms = append(rooms, &room)
	}

	parseLinks(rooms, links)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rooms
}
