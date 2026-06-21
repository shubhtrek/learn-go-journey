package main

import "fmt"

func safeDivision() {
	// Recover must be deferred
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Dividing by zero...")
	divisor := 0
	_ = 10 / divisor // Will trigger a panic
}

func main() {
	safeDivision()
	fmt.Println("Program successfully recovered and continues!")

	// 🎯 MINI CHALLENGE:
	// Write a nested defer block that prints messages to show execution order (First In, Last Out).
}
