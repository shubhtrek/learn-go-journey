package main

import (
	"fmt"
	"time"
)

func count(label string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", label, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Start counting concurrently
	go count("Goroutine")

	// Count on the main thread
	count("Main Thread")

	// Wait briefly to allow background routine to finish printing
	time.Sleep(500 * time.Millisecond)

	// 🎯 MINI CHALLENGE:
	// Try removing the time.Sleep at the end of the main function and run the code.
	// Why doesn't the Goroutine output print? Explain it to yourself!
}
