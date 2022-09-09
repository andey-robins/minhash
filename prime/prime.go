package prime

import (
	"math/big"
)

// IsPrime is a wrapper for the math/big function ProbablyPrime.
// the assumption there is that n will be <= 2^64, something
// assumed by this function. Values above that are not supported currently
// use ProbablyPrime itself
func IsPrime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}
