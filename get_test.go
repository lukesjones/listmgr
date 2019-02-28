package main

import "testing"

func TestGetAllLists(t *testing.T) {
	expected := "Here are all lists"
	output := getList("")

	if output != expected {
		t.Error("Output is different to expected!")
	}
}
func TestGetSpecificList(t *testing.T) {
	list := "test"
	expected := "Here is the list called: " + list

	output := getList(list)

	if output != expected {
		t.Error("Output is different to expected!")
	}
}
