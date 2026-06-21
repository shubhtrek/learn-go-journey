package main

import "fmt"

func main() {
	// Create an integer channel
	messages := make(chan string)

	// Spin up a goroutine to send a message
	go func() {
		messages <- "ping" // Sends message into channel
	}()

	// Read message from the channel (blocks until data arrives)
	msg := <-messages
	fmt.Println("Received:", msg)

	// 🎯 MINI CHALLENGE:
	// Create a channel of integers. Spin up a goroutine that sends the square
	// of 8 down the channel. Read and print it in the main function.
}
