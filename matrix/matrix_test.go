package matrix

import (
	"math"
	"testing"

	"github.com/andey-robins/minhash/helpers"
)

func TestCreateMatrix(t *testing.T) {
	rows := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix := NewMatrix()
	matrix.AddCol(rows[0])
	matrix.AddCol(rows[1])
	matrix.AddCol(rows[2])

	for i, r := range matrix.Cols {
		if !helpers.ListEqual(r, rows[i]) {
			t.Errorf("[Test %v] - Matrix row not as expected got=%v, exp=%v", i, r, rows[i])
		}
	}
}

func TestNewSigMatrix(t *testing.T) {
	sig := NewSigMatrix(3, 3)

	for i, row := range sig.Cols {
		for j, col := range row {
			if col != math.MaxInt64 {
				t.Errorf("[Test] - Element not equal to MaxInt at (%v, %v)", i, j)
			}
		}
	}
}

func TestGet(t *testing.T) {
	m := NewMatrix()
	m.AddCol([]int{1, 4, 7})
	m.AddCol([]int{2, 5, 8})
	m.AddCol([]int{3, 6, 9})

	tests := []struct {
		i   int
		j   int
		exp int
	}{
		{i: 0, j: 0, exp: 1},
		{i: 0, j: 1, exp: 2},
		{i: 0, j: 2, exp: 3},
		{i: 1, j: 0, exp: 4},
		{i: 1, j: 1, exp: 5},
		{i: 1, j: 2, exp: 6},
		{i: 2, j: 0, exp: 7},
		{i: 2, j: 1, exp: 8},
		{i: 2, j: 2, exp: 9},
	}

	for i, test := range tests {
		val := m.Get(test.i, test.j)
		if val != test.exp {
			t.Errorf("[Test %v] - Expected Get to return exp=%v, got=%v", i, test.exp, val)
		}
	}
}

func TestMatrixEqual(t *testing.T) {
	tests := []struct {
		m1  Matrix
		m2  Matrix
		exp bool
	}{
		{m1: *NewSigMatrix(2, 2), m2: *NewSigMatrix(2, 2), exp: true},
		{m1: *NewSigMatrix(2, 2), m2: *NewSigMatrix(2, 3), exp: false},
	}

	for i, test := range tests {
		if MatrixEqual(test.m1, test.m2) != test.exp {
			t.Errorf("[Test %v] - Matrices' equality was %v instead of %v", i, test.exp, !test.exp)
		}
	}
}
