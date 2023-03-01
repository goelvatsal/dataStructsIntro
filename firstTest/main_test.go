package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 1)
	if result != 3 {
		t.Fatalf("Result was incorrect, got: %d, want: %d.", result, 3)
	}
}
