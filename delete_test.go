package main

import "testing"

func TestDeleteNote(t *testing.T) {
	expected := "the delete function is not yet in use!"

	output := deleteNotes()

	if output != expected {
		t.Error("Output is different to expected!")
	}
}
