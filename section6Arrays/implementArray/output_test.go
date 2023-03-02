package main

import (
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	a := new(MyArray)

	a.push("hi")
	a.push("you")
	a.push("!")
	t.Log(*a)

	a.pop()
	//a.deleteAtIndex(0)
	t.Log(*a)

	a.push("are")
	a.push("nice")
	a.shiftItems(0)
	t.Log(*a)

	if a.length != 4 {
		t.Fatalf("Expected 4, got %d.", a.length)
	}

	if !reflect.DeepEqual(a.data, []string{"you", "are", "nice", "nice"}) {
		t.Fatalf(`Expected {"you", "are", "nice", "nice"}, got %s`, a.data)
	}
}
