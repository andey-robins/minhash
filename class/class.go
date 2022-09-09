package class

import (
	"github.com/andey-robins/minhash/matrix"
)

// ClassDriver will execute everything using the values provided
// in the class example. Useful for verifying all of the systems come
// together correctly
func ClassDriver() {
	s1 := []int{1, 0, 1, 0, 1}
	s2 := []int{0, 0, 1, 1, 0}
	s3 := []int{1, 1, 0, 1, 0}

	// h1 := func(x int) int { return (x + 1) % 5 }
	// h2 := func(x int) int { return (3*x + 1) % 5 }

	m := matrix.NewMatrix()
	m.AddRow(s1)
	m.AddRow(s2)
	m.AddRow(s3)

	approximateSig(5, 3)
}

func approximateSig(rows, cols int) {
	// sig := matrix.NewSigMatrix(rows, cols)
}
