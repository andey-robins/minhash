package matrix

import (
	"fmt"
	"math"

	"github.com/andey-robins/minhash/helpers"
)

// Matrix is a 2 dimensional data structure
type Matrix struct {
	Cols [][]int
}

// NewMatrix returns a pointer to a Matrix object
func NewMatrix() *Matrix {
	return &Matrix{}
}

// AddCol appends the argument col to the matrix
// it doesn't perform any length checking, allowing for variable
// width matrix rows
func (m *Matrix) AddCol(col []int) {
	m.Cols = append(m.Cols, col)
}

// Get will allow for (i,j) based indexing of the matrix
// returning the value at M_ij
func (m *Matrix) Get(i, j int) int {
	return m.Cols[j][i]
}

// Set will allow for (i,j) based indexing of the matrix
// to modify a value in the matrix
func (m *Matrix) Set(i, j, x int) {
	m.Cols[j][i] = x
}

func (m *Matrix) Print() {
	rows := make([][]int, len(m.Cols[0]))
	for i := range rows {
		rows[i] = make([]int, len(m.Cols))
	}
	// print rows
	for col := 0; col < len(m.Cols); col++ {
		for row := 0; row < len(rows); row++ {
			rows[row][col] = m.Get(row, col)
		}
	}

	for _, row := range rows {
		for _, v := range row {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
}

// NewSigMatrix returns a SIG matrix for use in the minhash algorithm
// this is a matrix of size rows by cols with all entries equal to
// the max int64 (the closest ints get to infinity)
func NewSigMatrix(rows, cols int) *Matrix {
	sig := NewMatrix()

	for i := 0; i < cols; i++ {
		col := []int{}
		for j := 0; j < rows; j++ {
			col = append(col, math.MaxInt64)
		}
		sig.AddCol(col)
	}

	return sig
}

func MatrixEqual(m1, m2 Matrix) bool {
	if len(m1.Cols) != len(m2.Cols) || len(m1.Cols[0]) != len(m2.Cols[0]) {
		return false
	}
	for i, col := range m1.Cols {
		if !helpers.ListEqual(col, m2.Cols[i]) {
			return false
		}
	}
	return true
}
