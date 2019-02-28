package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	err := createFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := os.Args[1]
	message := ""

	switch cmd {
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Please add list name.")
			os.Exit(1)
		}
		list := os.Args[2]

		message = createList(list)
	case "delete":
		message = deleteList()
	case "get":
		list := ""
		if len(os.Args) > 2 {
			list = os.Args[2]
		}
		message = getList(list)
	default:
		usage()
	}

	fmt.Println(message)
}

func usage() {
	usage := `Usage: %s [COMMAND]
Manage lists from your command line.

Available commands:
  create:  create a new list
  get:     get a list
  delete:  delete a specific list
`
	fmt.Printf(usage, os.Args[0])
	os.Exit(1)
}
