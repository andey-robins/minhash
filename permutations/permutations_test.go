package permutations

import (
	"testing"

	"github.com/andey-robins/minhash/helpers"
)

func TestPermuteList(t *testing.T) {
	tests := []struct {
		seed     int
		list     []int
		expected []int
	}{
		{seed: 1, list: []int{1, 2, 3}, expected: []int{1, 3, 2}},
		{seed: 2, list: []int{1, 2, 3}, expected: []int{2, 3, 1}},
		{seed: 1, list: []int{2, 2, 3}, expected: []int{2, 3, 2}},
		{seed: 5, list: []int{2, 4, 3, 1}, expected: []int{2, 3, 4, 1}},
	}

	for i, test := range tests {
		PermuteList(&test.list, test.seed)
		if !helpers.ListEqual(test.list, test.expected) {
			t.Errorf("[Test %v] - PermuteList failed for got=%v, exp=%v", i, test.list, test.expected)
		}
	}
}
