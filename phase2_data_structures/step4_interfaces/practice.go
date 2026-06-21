package main

// Shape interface defines behavior
type Shape interface {
	Area() float64
}

// TODO: Implement the Shape interface for Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Write Area method for Rectangle here


// TODO: Implement the Shape interface for Circle (use 3.14159 for pi)
type Circle struct {
	Radius float64
}

// Write Area method for Circle here
