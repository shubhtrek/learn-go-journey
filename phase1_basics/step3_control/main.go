package main

import "fmt"

func main() {
	// If-else block
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// For loop acting as a traditional loop
	fmt.Println("Loop 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 🎯 MINI CHALLENGE:
	// Write a loop from 1 to 30.
	// For multiples of 3, print "Chai".
	// For multiples of 5, print "Samosa".
	// For multiples of both 3 and 5, print "Chai-Samosa".
}
