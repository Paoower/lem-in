package main

import (
	"lem-in/src"
	"os"
)

func main() {
	if src.ArgsValidation() {
		src.ColonyMapping(os.Args[1])
	}
}
