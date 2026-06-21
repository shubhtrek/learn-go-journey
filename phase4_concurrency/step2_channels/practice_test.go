package main

import (
	"testing"
	"time"
)

func TestSendGreeting(t *testing.T) {
	ch := make(chan string)

	go SendGreeting(ch)

	select {
	case msg := <-ch:
		if msg != "Hello Channel" {
			t.Errorf("Expected 'Hello Channel', got %q", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Timed out waiting for channel response")
	}
}
