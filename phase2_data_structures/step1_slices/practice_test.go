package main

import (
	"reflect"
	"testing"
)

func TestFilterEven(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := []int{2, 4, 6, 8}

	got := FilterEven(input)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("FilterEven(%v) = %v; expected %v", input, got, expected)
	}
}
