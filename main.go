package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	cmd := os.Args[1]

	switch cmd {
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Please add note name.")
			os.Exit(1)
		}
		list := os.Args[2]
		createNote(list)
	case "list":
		listNotes()
	case "delete":
		deleteNotes()
	default:
		usage()
	}
}

func usage() {
	fmt.Println(`Usage: notemgr [OPTION]...
manage notes from your command line`)
	os.Exit(1)
}

func createNote(name string) {
	fmt.Println(name)
}
func listNotes() {
	fmt.Println("the list function is not yet in use!")
}
func deleteNotes() {
	fmt.Println("the delete function is not yet in use!")
}
