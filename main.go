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
	message := ""

	switch cmd {
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Please add note name.")
			os.Exit(1)
		}
		note := os.Args[2]

		message = createNote(note)
	case "list":
		message = listNotes()
	case "delete":
		message = deleteNotes()
	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Please add note name you wish me to get.")
			os.Exit(1)
		}
		note := os.Args[2]
		message = getNote(note)
	default:
		usage()
	}

	fmt.Println(message)
}

func usage() {
	fmt.Printf("Usage: %s [OPTION]...\nManage lists from your command line.\n", os.Args[0])
	os.Exit(1)
}
