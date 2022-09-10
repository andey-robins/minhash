package helpers

// ListEqual takes two slices that are comparable and
// returns true if the lists have the same elements in the
// same positions and are the same length
func ListEqual[T comparable](l1, l2 []T) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i, e := range l1 {
		if l2[i] != e {
			return false
		}
	}
	return true
}

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
