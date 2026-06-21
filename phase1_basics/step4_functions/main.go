package main

import "fmt"

// A simple function returning one value
func greet(name string) string {
	return "Hello " + name
}

// A function returning multiple values
func calculate(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	message := greet("Shubham")
	fmt.Println(message)

	total, difference := calculate(15, 5)
	fmt.Printf("Total: %d, Difference: %d\n", total, difference)

	// 🎯 MINI CHALLENGE:
	// Write a function `stats` that takes a slice/array of numbers (or just 3 numbers)
	// and returns their sum, average, and product.
}
