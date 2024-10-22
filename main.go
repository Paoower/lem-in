package main

import (
	"lem-in/src"
	"lem-in/src/farm"
	"os"
)

func main() {
	if src.ArgsValidation() {
		var f *farm.Farm = farm.NewFarm()
		f.Create(os.Args[1])
		f.SolveProblem()
	}
}
