package minhash

import (
	"github.com/andey-robins/minhash/helpers"
	"github.com/andey-robins/minhash/matrix"
)

// ApproximateSigMatrix takes an S and hash matrix `m` and approximates
// the sig matrix in place with the `sig` argument matrix
func ApproximateSigMatrix(m, sig *matrix.Matrix) {
	// this is an implementation of the Signature matrix approximation algorithm from lecture
	// 0. Initialize SIG(i, j) = infinity for all i, j
	//    (done outside this function with matrix.NewSigMatrix(...))
	// 1. For row i, compute h_1(i), h_2(i), ..., h_n(i)
	// 2. For each column j, if M(i, j) == 1, then for each k = 1, 2, ..., n
	//    set SIG(i, j) = min{ SIG(k, j), h_k(i) }
	for row := 0; row < len(m.Cols[0]); row++ {
		for col := 0; col < len(sig.Cols); col++ {
			if m.Get(row, col) == 1 {
				for hash := 0; hash < len(m.Cols)-len(sig.Cols); hash++ {
					h_ki := m.Get(row, len(sig.Cols)+hash)
					sig_kj := sig.Get(hash, col)
					sig.Set(hash, col, helpers.IntMin(sig_kj, h_ki))
				}
			}
		}
	}
}
