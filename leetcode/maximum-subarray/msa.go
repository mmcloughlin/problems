package msa

import "math"

func maxSubArray(nums []int) int {
	mx := math.MinInt64
	s := 0

	for _, num := range nums {
		s = max(s+num, num)
		mx = max(mx, s)
	}

	return mx
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
