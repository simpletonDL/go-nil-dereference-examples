package main

type A struct {
	p *int
}

var s, p *A

func _() {
	p = nil // TODO: naming
	s.p = nil
}

// Nil as zero value

// 1. Default initialization
// 2. Default field initialization

// 3. Map misssing lookup
// 4. Reading from a closed channel

// 5. Unmarshaling
