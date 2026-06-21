package main

import "fmt"

func main() {
	// 1. Explicit declaration
	var greeting string = "Hello World"
	fmt.Println(greeting)

	// 2. Short variable declaration (only inside functions!)
	count := 10
	fmt.Println("Count is:", count)

	// 3. Zero values demonstration
	var isBoiled bool
	var temperature float64
	var message string

	fmt.Printf("Zero values -> bool: %t, float: %f, string: %q\n", isBoiled, temperature, message)

	// 🎯 MINI CHALLENGE:
	// Declare an integer representing your target number of days to learn Go (e.g. 30).
	// Print it to the screen.
}
