package main

import "testing"

func TestRunSafely(t *testing.T) {
	panickingFunc := func() {
		panic("something went wrong")
	}

	safeFunc := func() {}

	res1 := RunSafely(panickingFunc)
	if res1 != "something went wrong" {
		t.Errorf("Expected 'something went wrong', got %q", res1)
	}

	res2 := RunSafely(safeFunc)
	if res2 != "success" {
		t.Errorf("Expected 'success', got %q", res2)
	}
}
