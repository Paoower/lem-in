package main

import (
	"lem-in/src"
	"os"
)

func main() {
	var rooms []*src.Room
	if src.ArgsValidation() {
		rooms = src.ColonyMapping(os.Args[1])
	}
	src.DisplayRooms(rooms)

}
