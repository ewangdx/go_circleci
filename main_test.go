package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 6
	actual := myAdd(3, 2)
	if (expected != actual) {
		t.Error("Nope")
	}
}
