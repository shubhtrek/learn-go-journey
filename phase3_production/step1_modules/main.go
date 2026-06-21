package main

import (
	"fmt"
	"step1_modules/mathutils"
)

func main() {
	res := mathutils.Add(10, 20)
	fmt.Println("10 + 20 =", res)

	// 🎯 MINI CHALLENGE:
	// 1. Look inside the mathutils directory and create a new file calculator.go inside it.
	// 2. Add an exported Subtract function, then call it here.
}
