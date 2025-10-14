package main

import "fmt"

func Copy(user *User) *User {
	return &User{
		Age:  user.Age,
		Name: user.Name,
	}
}

func HandleUser(user *User) {
	if user == nil {
		fmt.Println("User is nil")
	}
	copy := Copy(user)
	fmt.Printf("Copided, Name: %s, ", copy.Name)
}
