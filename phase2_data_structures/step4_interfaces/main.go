package main

import "fmt"

// Define the interface
type Speaker interface {
	Speak() string
}

type Human struct {
	Name string
}

func (h Human) Speak() string {
	return "Hello, my name is " + h.Name
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof! Woof!"
}

// Function accepting any Speaker
func Broadcast(s Speaker) {
	fmt.Println("Broadcast:", s.Speak())
}

func main() {
	person := Human{Name: "Shubham"}
	dog := Dog{}

	Broadcast(person)
	Broadcast(dog)

	// 🎯 MINI CHALLENGE:
	// Create an interface `Shape` with an `Area() float64` method.
	// Implement it with a `Circle` struct and a `Rectangle` struct.
	// Write a function to print the area of any Shape.
}
