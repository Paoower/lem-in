package main

import (
	"lem-in/src"
	"os"
)

func main() {
	var farm []*src.Room
	if src.ArgsValidation() {
		farm = src.FarmMapping(os.Args[1])
	}
	src.DisplayRooms(farm)
}
