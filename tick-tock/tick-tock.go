/*
	Tick Tock Clock
	Created: 11/18/19
*/

// Package declaration
package main

// Import format and time functions
import (
	"fmt"
	"time"
)

// Print introduction and show current time
func main() {
	fmt.Println("Welcome to the Go clock. Tick Tock.")

	fmt.Println("The time is:", time.Now())
}
