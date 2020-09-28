package main

import "fmt"

func fibSequence(n int) int {
	x := 0
	y := 1
	// We iterate until we get the desired position in the sequence
	for i := 0; i < n; i++ {
		// Using temporary variable to swap the variables
		temp := x
		x = y
		y = temp + x
	}
	return x
}

func main() {
	// Showing first 30 fibonacci numbers here
	for p := 0; p < 30; p++ {
		// Then we compute the numbers and display the result
		displayResult := fibSequence(p)
		fmt.Printf("Fibonacci number %v = %v\n", p, displayResult)
	}
}
