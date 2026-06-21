package main

import "sync"

type SafeStore struct {
	mutex sync.Mutex
	data  map[string]int
}

// TODO: Implement SafeStore.Set() to safely store a key-value pair using a Mutex.
func (s *SafeStore) Set(key string, val int) {
	// Write your code here
}

// TODO: Implement SafeStore.Get() to safely fetch a value.
func (s *SafeStore) Get(key string) int {
	// Write your code here
	return 0
}
