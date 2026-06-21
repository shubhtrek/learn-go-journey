package main

import "testing"

func TestGetPrice(t *testing.T) {
	menu := map[string]int{
		"Chai":   10,
		"Samosa": 15,
	}

	price, found := GetPrice(menu, "Chai")
	if price != 10 || !found {
		t.Errorf("Expected price 10 and found true, got price %d and found %t", price, found)
	}

	price, found = GetPrice(menu, "Jalebi")
	if price != -1 || found {
		t.Errorf("Expected price -1 and found false, got price %d and found %t", price, found)
	}
}
