package main

import "testing"

// TODO: Write unit tests for each function in functions.go
// Example:

func TestAdd(t *testing.T) {
	// Example test case for add function
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("add(2, 3) failed: expected %v, got %v", expected, result)
	}
}
