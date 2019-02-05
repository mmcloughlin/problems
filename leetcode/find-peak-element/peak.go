package peak

import "math"

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)
	for r-l > 1 {
		m := l + (r-l)/2

		if ispeak(nums, m) {
			return m
		}

		if lookup(nums, m-1) > lookup(nums, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func ispeak(nums []int, i int) bool {
	return lookup(nums, i-1) < lookup(nums, i) && lookup(nums, i) > lookup(nums, i+1)
}

func lookup(nums []int, i int) int {
	if i < 0 || i >= len(nums) {
		return math.MinInt64
	}
	return nums[i]
}
