package tools

import "log"

// Error helping function
func Check(e error) {
	log.SetFlags(0)
	if e != nil {
		log.Fatal("ERROR: ", e)
	}
}
