package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()

	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	numGoroutines := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final Counter Value (Should be 1000):", counter)
}
