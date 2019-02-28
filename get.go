package main

func getList(list string) string {
	if list == "" {
		return "Here are all notes"
	}
	return "Here is the list called: " + list
}
