package main

import (
	"fmt"
	"lem-in/src"
	"lem-in/src/farm"
	"lem-in/src/tools"
	"os"
)

func main() {
	tools.Check(src.ArgsValidation())
	var f *farm.Farm = farm.NewFarm()
	f.Create(os.Args[1])
	f.BFS()
	//f.PrintAllPaths()
	f.SortPaths()
	fmt.Print("\nPress Enter to start the visualization...")
	fmt.Scanln()
	f.VisualSolve()
}
