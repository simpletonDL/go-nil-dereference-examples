package main

type Config struct {
	State   string
	Options map[string]string
}

func main() {
	var c *Config
	c.State = "Utah"
	c.Options["Go"] = "West"
}
