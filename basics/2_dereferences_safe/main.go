package main

type User struct {
	Name string
}

func _() {
	var xs []int = nil
	_ = append(xs, 10)
	_ = len(xs)

	var m map[string]int = nil
	_ = m["123"]
	_ = len(m)

	for _, _ = range xs {
	}
	for _, _ = range m {
	}
}

// + save dereferences
