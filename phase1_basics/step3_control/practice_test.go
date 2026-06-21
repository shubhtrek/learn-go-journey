package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{3, "Chai"},
		{5, "Samosa"},
		{15, "ChaiSamosa"},
		{4, "4"},
		{9, "Chai"},
		{10, "Samosa"},
	}

	for _, tc := range tests {
		got := FizzBuzz(tc.input)
		if got != tc.expected {
			t.Errorf("FizzBuzz(%d) = %q; expected %q", tc.input, got, tc.expected)
		}
	}
}
