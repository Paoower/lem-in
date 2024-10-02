package main

import (
	"lem-in/src"
	"os"
)

func main() {
	var farm src.Farm
	if src.ArgsValidation() {
		farm = src.FarmMapping(os.Args[1])
	}
	src.DisplayFarm(farm)
}
