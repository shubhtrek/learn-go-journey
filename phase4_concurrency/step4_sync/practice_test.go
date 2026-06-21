package main

import (
	"sync"
	"testing"
)

func TestSafeStore(t *testing.T) {
	store := SafeStore{data: make(map[string]int)}

	var wg sync.WaitGroup
	numRoutines := 100

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			store.Set("key", val)
			_ = store.Get("key")
		}(i)
	}

	wg.Wait()
}
