package main

import (
	"fmt"
	"os"
)

func validateFile(data []byte) *ValidationError {
	return nil // lets consider it returns nil
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Error reading file")
	}
	err = validateFile(file)
	if err != nil {
		panic("Error is not nil! File is not valid")
	}
	fmt.Println("File is valid")
}
