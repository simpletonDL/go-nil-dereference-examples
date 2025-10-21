package main

type User struct {
	Name string
}

func _() {
	var xs []int = nil
	// ✅ Append to nil slice
	_ = append(xs, 10)
	// ✅ Get slice length
	_ = len(xs)

	var m map[string]int = nil
	// ✅ Read from nil map
	_ = m["123"]
	// ✅ Get map length
	_ = len(m)
	// ✅ Range over nil slice/map
	for _, _ = range xs {
	}
}

// + save dereferences
