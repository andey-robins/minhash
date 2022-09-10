package matrix

import "math"

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

// NewSigMatrix returns a SIG matrix for use in the minhash algorithm
// this is a matrix of size rows by cols with all entries equal to
// the max int64 (the closest ints get to infinity)
func NewSigMatrix(rows, cols int) *Matrix {
	sig := NewMatrix()

	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, math.MaxInt64)
		}
		sig.AddCol(row)
	}

	return sig
}
