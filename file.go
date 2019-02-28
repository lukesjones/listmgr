package main

import (
	"fmt"
	"os"
)

var (
	filename = "/home/luke/listmgr.txt"
)

func createFile() error {
	// get file info
	_, err := os.Stat(filename)

	// check if file exists
	exists := os.IsExist(err)
	if exists == true {
		return nil
	}

	// file does not exists at this point so create it
	_, err = os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating list file: %v", err)
	}

	return nil
}
