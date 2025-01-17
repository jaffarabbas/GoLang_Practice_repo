package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//create a function to read file name from a path in go
	listFiles("/home/darklord/Desktop/Github/")
}

// define the readFile function
func readFile(filePath string) {
	// open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create a buffer to read file
	buffer := make([]byte, 1024)

	// read file into buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", buffer[:n])
	}
}

// list all files in the directory
func listFiles(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name() + " " + file.Type().String())
		//print path
	}
}

func createDirectory(dirPath string) (string, error) {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", err
	}
	return dirPath, nil
}

func createFile(dirPath string) (string, error) {
	filePath := fmt.Sprintf("%s/newFile.txt", dirPath)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return filePath, nil
}

func deleteFile(dirPath string) (string, error) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println(err)
	}
	return "done", nil
}
