package farm

import (
	"errors"
	"lem-in/src/room"
	"strings"
)

// Fetches a room based on its name
func (farm *Farm) getRoom(name string) (*room.Room, error) {
	for _, room := range farm.Rooms {
		if room.Name == name {
			return room, nil
		}
	}
	return nil, errors.New("room not found")
}

// Checks if the ants' number is correct
func checkAnts(ants int) error {
	if ants <= 0 {
		return errors.New("invalid ants number")
	}
	return nil
}

// Checks if room was already saved
func (farm Farm) isRoomThere(room room.Room) bool {
	for _, r := range farm.Rooms {
		if r.Name == room.Name {
			return true
		}
	}
	return false
}

// Checks if a room's name is correct
func (farm *Farm) CheckRoomName(name string) error {
	if strings.HasPrefix(name, "L") {
		return errors.New("invalid room name")
	}
	return nil
}
