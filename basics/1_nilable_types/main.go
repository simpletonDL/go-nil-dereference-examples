package main

import "io"

var (
	p  *int           = nil // Pointers
	i  io.Reader      = nil // Interfaces
	s  []int          = nil // Slices
	m  map[string]int = nil // Maps
	ch chan int       = nil // Channels
	f  func()         = nil // Functions
)
