package main

import "testing"

// BasicFunction just adds two numbers
func BasicFunction(a, b int) int {
	return a + b
}

// TestBasicFunction is testing this basic function
func TestBasicFunction(t *testing.T) {
	answer := BasicFunction(5, 7)
	if answer != 12 {
		t.Error("unit test failed")
	}
}
