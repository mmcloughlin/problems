package solution

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	t := m + n
	steps := (t - 1) / 2
	l, r := 0, steps+1
	for r-l > 1 {
		m := l + (r-l)/2

		// at this step consider that the median is found by taking s1 steps into
		// nums1 and s2 steps into nums2
		s1, s2 := m, steps-m

		if index(nums1, s1-1) < index(nums2, s2) {
			l = m
		} else {
			r = m
		}
	}

	// what would the next element be?
	s1, s2 := l, steps-l
	x := [2]int{}
	for i := 0; i < 2; i++ {
		if index(nums1, s1) < index(nums2, s2) {
			x[i] = index(nums1, s1)
			s1++
		} else {
			x[i] = index(nums2, s2)
			s2++
		}
	}
	if t%2 == 1 {
		return float64(x[0])
	}
	return float64(x[0]+x[1]) / 2
}

func index(nums []int, i int) int {
	if i >= len(nums) {
		return math.MaxInt64
	}
	return nums[i]
}
