package main

import "testing"

func TestListNote(t *testing.T) {
	expected := "the list function is not yet in use!"
	output := listNotes()
	if output != expected {
		t.Error("Output is different to expected!")
	}
}
