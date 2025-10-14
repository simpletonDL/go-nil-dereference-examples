package main

import "fmt"

func main() {
	var ourCustomError *ValidationError = nil
	var err error = ourCustomError
	fmt.Printf("Error is nil: %t", err == nil)
}
