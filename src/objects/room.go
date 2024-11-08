package objects

import (
	"errors"
	"fmt"
	"lem-in/src/tools"
)

// Room datatype
type Room struct {
	Name     string
	X        int
	Y        int
	Ants     []*Ant
	Links    []*Room
}

// Checks if a room is reachable
func (room *Room) CheckRoomLink() error {
	if len(room.Links) == 0 {
		return errors.New("start or final room is unreachable")
	}
	return nil
}

// Display one single room
func (room Room) DisplayRoom() {
	fmt.Printf("Room: %s \n Links: ", room.Name)
	for _, link := range room.Links {
		fmt.Printf("%s ", link.Name)
	}
	tools.Ret()
}

// Room constructor
func NewRoom(name string, x int, y int) *Room {
	return &Room{
		Name:     name,
		X:        x,
		Y:        y,
		Ants:     []*Ant{},
		Links:    []*Room{},
	}
}
