package others

// 1. Default initialization and variables
// 2. Default initialization and fields
// 3. Default initialization of embedded fields

// Embeddings (interesting)
type Base struct {
	x int
}

type Derrived struct {
	*Base
}

//

func example1() {
	var m *Derrived
	_ = m.x
}

// Named type

type A struct {
	x int
}

type B *A

// Named type
