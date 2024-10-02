package main

import (
	"lem-in/src"
	"lem-in/src/farm"
	"os"
)

func main() {
	var f *farm.Farm = farm.NewFarm()
	if src.ArgsValidation() {
		f.Create(os.Args[1])
	}
	f.Display()
}
