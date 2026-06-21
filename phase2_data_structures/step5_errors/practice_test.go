package main

import "testing"

func TestSafeDivide(t *testing.T) {
	res, err := SafeDivide(10.0, 2.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res != 5.0 {
		t.Errorf("Expected 5.0, got %f", res)
	}

	_, err = SafeDivide(10.0, 0.0)
	if err == nil {
		t.Error("Expected error, got nil")
	} else if err.Error() != "cannot divide by zero" {
		t.Errorf("Expected 'cannot divide by zero', got %q", err.Error())
	}
}
