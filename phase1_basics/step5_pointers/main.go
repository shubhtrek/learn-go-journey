package main

import "fmt"

func updateValue(val int) {
	val = 100 // Modifies only the copy
}

func updateValuePointer(val *int) {
	*val = 100 // Modifies the actual value in memory
}

func main() {
	x := 10
	fmt.Println("Initial x:", x)

	updateValue(x)
	fmt.Println("After updateValue:", x)

	updateValuePointer(&x) // Pass the memory address of x
	fmt.Println("After updateValuePointer:", x)

	// 🎯 MINI CHALLENGE:
	// Write a function `swap` that takes two integer pointers (*int, *int)
	// and swaps the values of the variables they point to.
}
