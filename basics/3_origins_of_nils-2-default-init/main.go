package main

type User struct {
	Name string
}

func main() {
	var user *User    // user == nil
	user.Name = "123" // panic
}
