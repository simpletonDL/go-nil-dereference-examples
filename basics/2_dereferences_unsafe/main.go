package main

type User struct {
	Name string
}

func _() {
	var p *User = nil
	_ = *p     // ❌ Dereference of nil pointer
	_ = p.Name // ❌ Accessing field of nil pointer

	var e error = nil
	e.Error() // ❌ Calling method on nil interface

	var xs []User = nil
	_ = xs[0] // ❌ Indexing nil slice

	var m map[string]int = nil
	m["123"] = 10 // ❌ Assignment to nil map

	var f func() = nil
	f() // ❌ Calling nil function

	var ch chan int = nil
	ch <- 10 // 🔒 Send on nil channel
	<-ch     // 🔒 Receive from nil channel
}

// + save dereferences
