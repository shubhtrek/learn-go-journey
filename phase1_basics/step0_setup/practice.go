package main

import (
	"fmt"
)

// PrintSystemDiagnostics should print system and Go environment settings.
// Refer to the runtime and os packages in the official standard library docs.
// Expected outputs:
// 1. Go Version
// 2. Target OS and Architecture
// 3. Number of logical CPUs
// 4. Compiler Name
// 5. GOROOT and GOPATH environment variables
func PrintSystemDiagnostics() {
	// TODO: Replace the print statements below with actual runtime and os values
	fmt.Println("Go Version: [TODO]")
	fmt.Println("OS/Arch: [TODO]")
	fmt.Println("Logical CPUs: [TODO]")
	fmt.Println("Compiler: [TODO]")
	fmt.Println("GOROOT: [TODO]")
	fmt.Println("GOPATH: [TODO]")

	// Hint:
	// - Use runtime.Version() for Go version
	// - Use runtime.GOOS and runtime.GOARCH for OS/Arch
	// - Use runtime.NumCPU() for CPUs
	// - Use runtime.Compiler for the compiler
	// - Use os.Getenv("GOROOT") (or runtime.GOROOT()) and os.Getenv("GOPATH") for variables
}
