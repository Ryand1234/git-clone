package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program name:", os.Args[0])

	if len(os.Args) == 1 {
		Usage()
		return
	}
	supportedCommands := []string{"init", "help", "add"}

	command := os.Args[1]
	found := false
	for _, item := range supportedCommands {
		if item == command {
			found = true
			break
		}
	}

	if !found {
		CommandNotFound(command)
		return
	}

	if command == "help" {
		Usage()
		return
	}

	if command == "init" {
		if len(os.Args) <= 2 {
			fmt.Println("Error: Please specify repository name")
			return
		}
		dirName := os.Args[2]
		Init(dirName)
		return
	}

	if command == "add" {
		if len(os.Args) <= 2 {
			fmt.Println("Error: Please specify repository name")
			return
		}
		for _, fileName := range os.Args[2:] {
			Add(fileName)
		}
		return
	}
}
