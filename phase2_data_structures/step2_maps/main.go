package main

import "fmt"

func main() {
	// Initialize a map
	menu := make(map[string]int)
	menu["Samosa"] = 15
	menu["Chai"] = 10
	menu["Kachori"] = 20

	fmt.Println("Menu:", menu)

	// Checking if key exists (comma-ok idiom)
	price, exists := menu["Jalebi"]
	if exists {
		fmt.Println("Jalebi is on the menu for:", price)
	} else {
		fmt.Println("❌ Jalebi is not available today!")
	}

	// 🎯 MINI CHALLENGE:
	// Create a map that tracks student names (keys) and their exam scores (values).
	// Add 3 students, delete 1, and print the remaining list.
}
