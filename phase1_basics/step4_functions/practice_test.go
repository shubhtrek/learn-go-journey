package main

import "testing"

func TestCalculateRectangle(t *testing.T) {
	area, perimeter := CalculateRectangle(5.0, 10.0)

	if area != 50.0 {
		t.Errorf("Expected area 50.0, got %f", area)
	}
	if perimeter != 30.0 {
		t.Errorf("Expected perimeter 30.0, got %f", perimeter)
	}
}
