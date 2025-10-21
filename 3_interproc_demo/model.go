package main

import (
	"errors"
)

type User struct {
	Age  int
	Name string
}

func NewUser(age int, name string) *User {
	user, err := CreateUser(age, name)
	if err != nil {
		return nil
	}
	return user
}

func CreateUser(age int, name string) (*User, error) {
	if !isValidName(name) {
		return nil, errors.New("invalid name")
	}
	if !isValidAge(age) {
		return nil, nil
	}
	return &User{Age: age, Name: name}, nil
}
