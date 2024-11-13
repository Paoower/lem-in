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
	f.Create(os.Args[1])
	f.BFS()
	f.PrintAllPaths()
	f.SortPaths()
}
