package main

import (
	"fmt"
	"lem-in/src"
	"lem-in/src/farm"
	"os"
)

func main() {
	if src.ArgsValidation() {
		var f *farm.Farm = farm.NewFarm()
		f.Create(os.Args[1])
		f.GetAllPaths()
		f.SortPaths()
		for _, path := range f.PathObject.Paths {
			fmt.Println(path.Rooms[1])
		}
		fmt.Println(f.PathObject.PathsCap)
	}
}
