package minhash

import (
	"testing"

	"github.com/andey-robins/minhash/matrix"
)

func TestApproximateSigMatrix(t *testing.T) {
	s1 := []int{1, 0, 1, 0, 1}
	s2 := []int{0, 0, 1, 1, 0}
	s3 := []int{1, 1, 0, 1, 0}

	h1 := func(x int) int { return (x + 1) % 5 }
	h2 := func(x int) int { return (3*x + 1) % 5 }

	m := matrix.NewMatrix()
	m.AddCol(s1)
	m.AddCol(s2)
	m.AddCol(s3)

	sig := matrix.NewSigMatrix(2, 3)

	h1Row := []int{}
	h2Row := []int{}

	for i := 0; i < len(m.Cols[0]); i++ {
		h1Row = append(h1Row, h1(i))
		h2Row = append(h2Row, h2(i))
	}

	m.AddCol(h1Row)
	m.AddCol(h2Row)

	ApproximateSigMatrix(m, sig)

	expectedSig := matrix.NewMatrix()
	expectedSig.AddCol([]int{0, 1})
	expectedSig.AddCol([]int{3, 0})
	expectedSig.AddCol([]int{1, 0})

	if !matrix.MatrixEqual(*expectedSig, *sig) {
		t.Errorf("[Test 1] - Sig was not equal to expected sig.")
	}
}
