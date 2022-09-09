package helpers

import "testing"

func TestListEqual(t *testing.T) {
	tests := []struct {
		l1  []int
		l2  []int
		exp bool
	}{
		{l1: []int{1, 2, 3}, l2: []int{1, 2, 3}, exp: true},
		{l1: []int{1, 3, 3}, l2: []int{1, 3, 3}, exp: true},
		{l1: []int{1, 2, 3, 4}, l2: []int{1, 2, 3, 4}, exp: true},
		{l1: []int{1, 2}, l2: []int{1, 2}, exp: true},
		{l1: []int{3, 3, 3}, l2: []int{3, 3, 3}, exp: true},
		{l1: []int{}, l2: []int{}, exp: true},
		{l1: []int{1, 2, 3, 4}, l2: []int{1, 2, 3}, exp: false},
		{l1: []int{}, l2: []int{1}, exp: false},
		{l1: []int{1, 2, 3, 4}, l2: []int{1, 2, 34}, exp: false},
		{l1: []int{1, 2}, l2: []int{2, 1}, exp: false},
		{l1: []int{3, 3, 3}, l2: []int{3, 3, 3, 3}, exp: false},
		{l1: []int{1, 2}, l2: []int{}, exp: false},
	}

	for i, test := range tests {
		if ListEqual(test.l1, test.l2) != test.exp {
			t.Errorf("[Test %v] - Lists l1=%v and l2=%v equality should be %v but isn't", i, test.l1, test.l2, test.exp)
		}
	}
}
