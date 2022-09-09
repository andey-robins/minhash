package permutations

import "math/rand"

// PermuteList takes the elements in the argument list and shuffles
// them within the array, in place. The seed is used to seed the random
// number generator to provide an avenue for deterministic testing/outcomes
func PermuteList(list *[]int, seed int) {
	rand.Seed(int64(seed))
	rand.Shuffle(len(*list), func(i, j int) { (*list)[i], (*list)[j] = (*list)[j], (*list)[i] })
}
