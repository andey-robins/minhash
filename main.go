package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/andey-robins/minhash/class"
	"github.com/andey-robins/minhash/jaccard"
	"github.com/andey-robins/minhash/language"
	"github.com/andey-robins/minhash/matrix"
	"github.com/andey-robins/minhash/minhash"
	"github.com/andey-robins/minhash/prime"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information.")
	}

	var fPrime, seed, subsets int
	var help, classOnly bool
	flag.IntVar(&fPrime, "L", 11, "the value of L in S = {0, 1, 2, 3, ... L-1}")
	flag.IntVar(&subsets, "Q", 10, "the number of subsets of S")
	flag.IntVar(&seed, "seed", int(time.Now().UnixNano()), "the seed for random number generation")
	flag.BoolVar(&help, "help", false, "print help information")
	flag.BoolVar(&classOnly, "class", false, "set values to default for class example")
	flag.Parse()

	if help {
		pad := func() {
			fmt.Printf("\n\n")
		}
		pad()
		fmt.Println(" Welcome to the minhash explorer tool!")
		fmt.Print(" This code is licensed under GPLv3")
		pad()
		fmt.Println(" Args:")
		fmt.Println("  -L        The upper bound in defining S = {0, 1, 2, ..., L-1}")
		fmt.Println("  -Q        The number of subsets of S to create")
		fmt.Println("  -seed     The seed for all RNG used in this code. Defaults to Unix Time if none provided")
		fmt.Println("  -class    Automatically set defaults to the class example")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if classOnly {
		class.ClassDriver()
		return
	}

	if !prime.IsPrime(fPrime) {
		fmt.Println("Argument -L is not prime.")
		return
	}

	rand.Seed(int64(seed))
	driver(fPrime, subsets)
}

func driver(l, q int) {
	fmt.Printf("Running program with values L=%v, Q=%v\n\n", l, q)

	// Step 1 from homework
	lang := language.New[int]()
	lang.AddIntRange(0, l)
	fmt.Println("Generated sparse columns")

	// Step 2 from homework
	sparseMatrix := matrix.NewMatrix()

	for i := 0; i < q; i++ {
		sparseMatrix.AddCol(lang.SparseList())
	}

	fmt.Println("Constructed sparse matrix")

	// Step 3 from homework
	// calculate n to be a prime number less than L
	n := l - 1
	for !prime.IsPrime(n) {
		n--
	}

	// call hashbuilder with the prime number and size to generate a hash
	// function for some given row number x
	hashBuilder := func(prime, l int) func(x int) int {
		return func(x int) int {
			return (prime*x + 2) % l
		}
	}

	// compute the hashes and add them to the column matrix
	for i := 1; i <= n; i++ {
		hash := hashBuilder(i, l)
		hashRow := []int{}

		for i := 0; i < len(sparseMatrix.Cols[0]); i++ {
			hashRow = append(hashRow, hash(i))
		}

		sparseMatrix.AddCol(hashRow)
	}

	fmt.Println("Constructed hashes")

	sig := matrix.NewSigMatrix(n, q)
	minhash.ApproximateSigMatrix(sparseMatrix, sig)

	left := rand.Intn(len(sig.Cols))
	right := rand.Intn(len(sig.Cols))

	fmt.Printf("Approximated Jaccard = %v\n", jaccard.JaccardSimilarity(sig.Cols[left], sig.Cols[right]))
	fmt.Printf("True Jaccard         = %v\n", jaccard.JaccardSimilarity(sparseMatrix.Cols[left], sparseMatrix.Cols[right]))
}
