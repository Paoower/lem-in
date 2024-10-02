package farm

import (
	"fmt"
	"lem-in/src/tools"
)

// Display all the rooms and their links based on a given slice
func (farm Farm) displayRooms() {
	// Print all rooms
	for i := 0; i < len(farm.Rooms); i++ {
		// Display first room
		if i == 0 {
			fmt.Println("First room: ")
			farm.Rooms[0].DisplayRoom()
			tools.Ret()
			continue
		}

		// Display final room
		if i == len(farm.Rooms)-1 {
			tools.Ret()
			fmt.Println("Final room: ")
			farm.Rooms[len(farm.Rooms)-1].DisplayRoom()
			continue
		}
		farm.Rooms[i].DisplayRoom()
	}
}

// Display the whole farm
func (farm Farm) Display() {
	fmt.Println("Number of ants: ", farm.Ants)
	tools.Ret()
	farm.displayRooms()
}
