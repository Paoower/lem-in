package src

import (
	"errors"
	"os"
)

// Is the path given correct ?
func isARealFile(txtfilepath string) bool {
	_, error := os.Stat(txtfilepath)
	return errors.Is(error, os.ErrNotExist)
}

// Was the application properly launched with a valid argument ?
func ArgsValidation() error {
	if len(os.Args) != 2 {
		return errors.New("invalid argument")
	} else {
		txtfilepath := os.Args[1]
		if isARealFile(txtfilepath) {
			return errors.New("the specified file wasn't found")
		} else {
			//Good number of argument, and the txt file exist at the end of the given path
			return nil
		}
	}
}
