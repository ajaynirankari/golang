package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	message := "Message Hello World !"
	println(message)

	fileName := "test.txt"

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		println("Error getting file info with error: ", err)
	}

	fmt.Println("File info: ", fileInfo, "IsExist:", os.IsExist(err), "IsNotExist:", os.IsNotExist(err))

	if os.IsNotExist(err) {
		println("--------------- File does not exist-------------")

		file, err := os.Create(fileName)
		if err != nil {
			println("Error creating file with error: ", err)
			return
		}

		for i := 0; i < 10; i++ {
			message := fmt.Sprintf("Hello, World_[%d] at time: %v\n", i+1, time.Now())
			file.WriteString(message)
		}
		println("File created successfully")
		file.Close()
	}

	// read file
	file, err := os.Open(fileName)
	if err != nil {
		println("Error opening file with error: ", err)
		return
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		println("Error reading file with error: ", err)
		return
	}
	fileContent := string(data)
	println("Read data from file:\n", fileContent)
	file.Close()

	os.Remove(fileName)

	// create temp file
	tempFile, err := os.CreateTemp("", "tempFile_*.txt")
	if err != nil {
		println("Error creating temp file with error: ", err)
		return
	}
	tempFile.WriteString("Hello, World! from temp file")
	tempFile.Close()
	os.Remove(tempFile.Name())

	// create directory
	dirName := "testDir"
	os.Mkdir(dirName, 0755)
	os.Remove(dirName)
}
