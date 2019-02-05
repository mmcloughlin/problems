package maxprod

import "math"

func maxProduct(nums []int) int {
	rmn, rmx := 1, 1
	mx := math.MinInt64

	for _, num := range nums {
		rmn, rmx = min(rmn*num, rmx*num, num), max(rmn*num, rmx*num, num)
		mx = max(mx, rmx)
	}

	return mx
}

func max(xs ...int) int {
	m := math.MinInt64
	for _, x := range xs {
		if x > m {
			m = x
		}
	}
	return m
}

func min(xs ...int) int {
	m := math.MaxInt64
	for _, x := range xs {
		if x < m {
			m = x
		}
	}
	return m
}
