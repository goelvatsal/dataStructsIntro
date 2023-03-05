package main

import "testing"

func TestReverseString(t *testing.T) {
	if reverseString("Hello, I am Vatsal.") != ".lastaV ma I ,olleH" {
		t.Logf("Incorrect output, got %q and expected %q", reverseString("Hello, I am Vatsal."), ".lastaV ma I ,olleH")
	}
}
