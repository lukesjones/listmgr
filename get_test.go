package main

import "testing"

func TestGetAllNotes(t *testing.T) {
	expected := "Here are all notes"
	output := getList("")

	if output != expected {
		t.Error("Output is different to expected!")
	}
}
func TestGetSpecificNote(t *testing.T) {
	list := "test"
	expected := "Here is the list called: " + list

	output := getList(list)

	if output != expected {
		t.Error("Output is different to expected!")
	}
}
