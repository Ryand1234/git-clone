package main

import (
	"fmt"
	"os"
	"sync"
)

func createConfig(location string) {
	fileName := location + "/config"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	content := []byte("[core]\nrepositoryformatversion = 0\nfilemode = false\nbare = false\nlogallrefupdates = true\nsymlinks = false\nignorecase = true")
	_, writeErr := file.Write(content)
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		return
	}

	defer file.Close()
}

func createHead(location string) {
	fileName := location + "/HEAD"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	content := []byte("ref: refs/heads/master")
	_, writeErr := file.Write(content)
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		return
	}
	defer file.Close()
}

func createDescription(location string) {
	fileName := location + "/description"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	content := []byte("Unnamed repository; edit this file 'description' to name the repository.")
	_, writeErr := file.Write(content)
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		return
	}
	defer file.Close()
}

func createConfigFiles(location string) {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		createConfig(location)
		wg.Done()
	}()

	go func() {
		createDescription(location)
		wg.Done()
	}()

	go func() {
		createHead(location)
		wg.Done()
	}()

	wg.Wait()
}

func createConfigDirectories(location string) {
	directoriesToCreate := []string{"hooks", "info", "objects", "refs"}

	for _, item := range directoriesToCreate {
		dirName := location + "/" + item
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}
}

func Init(dirName string) {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	location := dirName + "/" + ".git"
	err = os.Mkdir(location, 0755)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Creating config files.")
	createConfigFiles(location)
	createConfigDirectories(location)
}
