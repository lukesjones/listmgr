package main

import "testing"

func TestCreateNote(t *testing.T) {
	expected := "You created a note called: "
	note := "test"

	output := createNote(note)

	if output != expected {
		t.Error("Output is different to expected!")
	}
}
