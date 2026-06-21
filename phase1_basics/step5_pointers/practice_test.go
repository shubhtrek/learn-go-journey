package main

import "testing"

func TestDoubleValue(t *testing.T) {
	x := 10
	DoubleValue(&x)
	if x != 20 {
		t.Errorf("Expected x to be 20, got %d", x)
	}
}

func TestSwapValues(t *testing.T) {
	a, b := 5, 10
	SwapValues(&a, &b)
	if a != 10 || b != 5 {
		t.Errorf("Expected a=10, b=5; got a=%d, b=%d", a, b)
	}
}
