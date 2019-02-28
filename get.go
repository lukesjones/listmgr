package main

func getList(list string) string {
	if list == "" {
		return "Here are all lists"
	}
	return "Here is the list called: " + list
}
