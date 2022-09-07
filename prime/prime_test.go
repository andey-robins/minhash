package prime

import "testing"

func TestIsPrime(t *testing.T) {
	numbers := []struct {
		num   int
		prime bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
	}

	for i, test := range numbers {
		isPrime := IsPrime(test.num)
		if isPrime != test.prime {
			t.Errorf("[Test %d] - IsPrime failed for num=%v, exp=%v, got=%v", i, test.num, test.prime, isPrime)
		}
	}
}
