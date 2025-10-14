package main

type Config struct {
	State   string
	Options map[string]string
}

func main() {
	c := Config{
		State: "Utah",
	}
	// c.Options == nil
	c.Options["Go"] = "West" // panic
}
