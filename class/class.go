package class

import (
	"fmt"

	"github.com/andey-robins/minhash/jaccard"
	"github.com/andey-robins/minhash/matrix"
	"github.com/andey-robins/minhash/minhash"
)

// ClassDriver will execute everything using the values provided
// in the class example. Useful for verifying all of the systems come
// together correctly. Also included in `go test ./...` so if the tests
// pass, then the minhash algorithm components are working
func ClassDriver() {
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

	minhash.ApproximateSigMatrix(m, sig)

	fmt.Println("\nExpected Matrix:")
	fmt.Println("0 3 1")
	fmt.Println("1 0 0")
	fmt.Println("\nCalculated Matrix:")
	sig.Print()
	fmt.Println()

	expectedSig := matrix.NewMatrix()
	expectedSig.AddCol([]int{0, 1})
	expectedSig.AddCol([]int{3, 0})
	expectedSig.AddCol([]int{1, 0})

	if matrix.MatrixEqual(*sig, *expectedSig) {
		fmt.Println("Matrices were equal. Everything worked!")
	} else {
		fmt.Println("Matrices were not equal. Something went wrong.")
	}
	fmt.Println()

	trueSimilarity := jaccard.JaccardSimilarity(m.Cols[0], m.Cols[2])
	estimatedSimilarity := jaccard.JaccardSimilarity(sig.Cols[0], sig.Cols[2])

	fmt.Printf("The true Jaccard similarity is %v while the estimated similarity is %v\n\n", trueSimilarity, estimatedSimilarity)

}
