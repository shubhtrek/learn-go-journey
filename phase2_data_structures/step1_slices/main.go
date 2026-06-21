package main

import "fmt"

func main() {
	// Arrays: fixed size
	var arr [3]string = [3]string{"Chai", "Coffee", "Water"}
	fmt.Println("Array:", arr)

	// Slices: dynamic size
	slice := []string{"Lassi", "Buttermilk"}
	fmt.Println("Initial Slice:", slice)

	// Append elements
	slice = append(slice, "Lemonade")
	fmt.Println("After Append:", slice)

	// Slicing an existing slice/array
	subSlice := slice[1:3] // index 1 and 2
	fmt.Println("Sub-slice:", subSlice)

	// 🎯 MINI CHALLENGE:
	// Create a slice of your 5 favorite programming languages.
	// Print it. Slice it to get the top 3. Append a new language to the sub-slice and check what happens.
}
