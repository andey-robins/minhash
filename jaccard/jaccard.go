package jaccard

// JaccardSimilarity will compute the JaccardSimilarity for
// two lists
func JaccardSimilarity(l1, l2 []int) float64 {
	intersect := 0
	union := 0
	for i, e := range l1 {
		if e == l2[i] && e == 1 {
			intersect++
		}

		if e == 1 || l2[i] == 1 {
			union++
		}
	}

	if intersect == 0 && union == 0 {
		return 0
	}

	return float64(intersect) / float64(union)
}
