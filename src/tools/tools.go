package tools

import (
	"fmt"
	"time"
)

// Prints a new line
func Ret() {
	fmt.Print("\n")
}

// Debug message with 1 second pause
func Debug(message string) {
	fmt.Println(message)
	time.Sleep(1 * time.Second)
}
