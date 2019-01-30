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
		note := os.Args[2]
		createNote(note)
	case "list":
		listNotes()
	case "delete":
		deleteNotes()
	default:
		usage()
	}
}

func usage() {
	fmt.Printf("Usage: %s [OPTION]...\nManage lists from your command line.\n", os.Args[0])
	os.Exit(1)
}
