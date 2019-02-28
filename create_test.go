package main

import "testing"

func TestCreateList(t *testing.T) {
	expected := "You created a note called: test"
	list := "test"

	output := createList(list)

	if output != expected {
		t.Error("Output is different to expected!")
	}
}
