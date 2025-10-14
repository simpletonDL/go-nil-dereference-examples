package main

type User struct {
	Name     *string           `json:"name"`
	Email    string            `json:"email"`
	Tags     []string          `json:"tags"`
	Settings map[string]string `json:"settings"`
}

func main() {
	//NIL
	var x *int
	println(*x)
}
