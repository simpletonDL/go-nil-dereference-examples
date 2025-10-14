package main

import (
	config "examples/1_var_shadowing/utils"
	"fmt"
)

// Spoken: Hidden issue, works well for the most cases and for the most time.

func main() {
	conf := config.LoadFromFile()
	if conf == nil {
		conf, err := config.LoadFromEnvironment()
		if conf == nil || err != nil {
			return
		}
	}
	fmt.Printf("Host: %s, Port: %s", conf.Host, conf.Port)
}
