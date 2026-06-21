package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}
	fmt.Println("Result:", res)

	// 🎯 MINI CHALLENGE:
	// Write a function `checkAge` that returns an error if age is less than 18.
	// Call it from main and print appropriate messages.
}
