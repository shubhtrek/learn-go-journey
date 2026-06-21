package main

import (
	"math"
	"testing"
)

func TestInterfaces(t *testing.T) {
	var s Shape

	s = Rectangle{Width: 4, Height: 5}
	if s.Area() != 20.0 {
		t.Errorf("Expected rectangle area 20.0, got %f", s.Area())
	}

	s = Circle{Radius: 10}
	expectedCircleArea := 3.14159 * 10 * 10
	if math.Abs(s.Area()-expectedCircleArea) > 0.001 {
		t.Errorf("Expected circle area %f, got %f", expectedCircleArea, s.Area())
	}
}
