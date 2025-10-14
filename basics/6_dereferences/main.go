package __dereferences

type User struct {
	Name string
}

func _() {
	var p *User = nil
	_ = *p
	_ = p.Name

	var e error = nil
	e.Error()

	var xs []User = nil
	_ = xs[0]

	var m map[string]int = nil
	m["123"] = 10
}

// + save dereferences
