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
	matrix.AddRow(rows[0])
	matrix.AddRow(rows[1])
	matrix.AddRow(rows[2])

	for i, r := range matrix.Rows {
		if !helpers.ListEqual(r, rows[i]) {
			t.Errorf("[Test %v] - Matrix row not as expected got=%v, exp=%v", i, r, rows[i])
		}
	}
}

func TestNewSigMatrix(t *testing.T) {
	sig := NewSigMatrix(3, 3)

	for i, row := range sig.Rows {
		for j, col := range row {
			if col != math.MaxInt64 {
				t.Errorf("[Test] - Element not equal to MaxInt at (%v, %v)", i, j)
			}
		}
	}
}
