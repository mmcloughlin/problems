package solution

func intersect(nums1 []int, nums2 []int) []int {
	count := map[int]int{}

	for _, n := range nums1 {
		count[n]++
	}

	intersection := []int{}
	for _, n := range nums2 {
		if count[n] > 0 {
			intersection = append(intersection, n)
			count[n]--
		}
	}

	return intersection
}
