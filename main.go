package main

import (
	"lem-in/src"
	"lem-in/src/farm"
	"lem-in/src/tools"
	"os"
)

func main() {
	tools.Check(src.ArgsValidation())
	var f *farm.Farm = farm.NewFarm()
	if !f.Create(os.Args[1]) {
		return
	}
	f.BFS()
	//f.PrintAllPaths()
	f.SortPaths()
	// fmt.Print("\nPress Enter to start the visualization...")
	// fmt.Scanln()
	f.VisualSolve()
}
