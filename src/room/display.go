package room

import (
	"fmt"
	"lem-in/src/tools"
)

// Display one single room
func (room Room) DisplayRoom() {
	fmt.Printf("Room: %s \n Links: ", room.Name)
	for _, link := range room.Links {
		fmt.Printf("%s ", link.Name)
	}
	tools.Ret()
}
