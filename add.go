package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func getCurrentDirectory() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return currentDir, nil
}

func createIndexFile(location string) {
	fileName := location + "/.git/index"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
}

func getFileHash(fileLocation string) (string, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()

	_, copyErr := io.Copy(hash, file)
	if copyErr != nil {
		return "", copyErr
	}

	hashByte := hash.Sum(nil)
	hashHex := fmt.Sprintf("%x", hashByte)

	return hashHex, nil
}

func Add(location string) {
	fmt.Println("Adding file ", location)
	dirName, err := getCurrentDirectory()
	if err != nil {
		fmt.Println(err)
		return
	}
	indexFile, err := os.Open(".git/index")
	if err != nil {
		createIndexFile(dirName)
		indexFile, err = os.Open(".git/index")
		if err != nil {
			return
		}
	}
	defer indexFile.Close()
	fileLocation := dirName + "/" + location
	fileHash, err := getFileHash(fileLocation)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file Hash: ", fileHash)
	// addFileHashToIndex(fileHash)
}
