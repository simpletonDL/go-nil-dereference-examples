package main

import (
	"fmt"
)

func main() {
	user := NewUser(0, "Sergey")
	fmt.Printf("Name: %s, ", user.Name)
}
