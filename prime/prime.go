package prime

import (
	"math/big"
)

// IthPrime is a structure used to provide memoized values of prime
// numbers
type IthPrime struct {
	primes []int
}

// IsPrime is a wrapper for the math/big function ProbablyPrime.
// the assumption there is that n will be <= 2^64, something
// assumed by this function. Values above that are not supported currently
// use ProbablyPrime itself
func IsPrime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}

// NewIthPrime will create an IthPrime object for prime number memoization
func NewIthPrime() *IthPrime {
	return &IthPrime{
		primes: []int{1, 2, 3, 5, 7, 11},
	}
}

// GetPrime will return the ith prime number, 0 indexed
// this function also handles memoizing primes via the IthPrime
// struct declared in this file
func (ip *IthPrime) GetPrime(i int) int {
	if i < len(ip.primes) {
		return ip.primes[i]
	}

	nextPrime := ip.primes[len(ip.primes)-1] + 1

	for !IsPrime(nextPrime) {
		nextPrime++
	}

	ip.primes = append(ip.primes, nextPrime)

	return ip.GetPrime(i)
}
