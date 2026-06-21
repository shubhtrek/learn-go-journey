package main

import "testing"

func TestGetVariables(t *testing.T) {
	i, f, s, b := GetVariables()

	if i != 42 {
		t.Errorf("Expected 42, got %d", i)
	}
	if f != 3.14 {
		t.Errorf("Expected 3.14, got %f", f)
	}
	if s != "Go is awesome" {
		t.Errorf("Expected 'Go is awesome', got %q", s)
	}
	if b != true {
		t.Errorf("Expected true, got %t", b)
	}
}
