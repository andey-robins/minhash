package jaccard

// JaccardSimilarity will compute the JaccardSimilarity for
// two lists
func JaccardSimilarity[T comparable](l1, l2 []T) float64 {
	similar := 0
	for i, e := range l1 {
		if e == l2[i] {
			similar++
		}
	}

	return float64(similar) / float64(len(l1))
}
