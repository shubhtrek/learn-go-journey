package main

import "testing"

func TestReverse(t *testing.T) {
	// Table-driven tests!
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"Go", "oG"},
		{"", ""},
		{"tea", "aet"},
	}

	for _, tc := range tests {
		result := Reverse(tc.input)
		if result != tc.expected {
			t.Errorf("Reverse(%q) = %q; expected %q", tc.input, result, tc.expected)
		}
	}
}
