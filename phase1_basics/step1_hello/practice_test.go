package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintDetails(t *testing.T) {
	// Redirect stdout to capture printed text
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintDetails("Shubham", "🍕")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Name: Shubham") {
		t.Errorf("Expected output to contain 'Name: Shubham', got: %q", output)
	}
	if !strings.Contains(output, "Favorite Snack: 🍕") {
		t.Errorf("Expected output to contain 'Favorite Snack: 🍕', got: %q", output)
	}
}
