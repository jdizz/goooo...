/*
	Big Number in Go
	Created 11/19/19
*/

// Package declaration
package main

// Import needed packages
import (
	"fmt"
	"math/rand"
)

// How fast can we print a huuuggggeeee number?
// Too big overflows the int...hmmm
func main() {
	fmt.Println("My favorite really big number is:",
		rand.Intn(1000000000000000000))
}
