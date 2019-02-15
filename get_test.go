package main

import "testing"

func TestGetNote(t *testing.T) {
	expected := "Here is the list called: test"
	note := "test"

	output := getNote(note)

	if output != expected {
		t.Error("Output is different to expected!")
	}
}
