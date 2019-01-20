package threesum

import "sort"

func threeSum(nums []int) [][]int {
	triplets := [][]int{}

	// Count of each number.
	count := map[int]int{}
	for _, n := range nums {
		count[n]++
	}

	// Unique numbers (sorted).
	N := len(count)
	uniq := make([]int, 0, N)
	for n := range count {
		uniq = append(uniq, n)
	}
	sort.Ints(uniq)

	for i := 0; i < N; i++ {
		b := uniq[i]
		count[b]--
		for j := 0; j <= i; j++ {
			a := uniq[j]
			if count[a] == 0 {
				continue
			}
			count[a]--
			c := -(a + b)
			if c >= b && count[c] > 0 {
				triplet := []int{a, b, c}
				triplets = append(triplets, triplet)
			}
			count[a]++
		}
		count[b]++
	}

	return triplets
}
