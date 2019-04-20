package solution

import "math"

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// Cumulative min from the right.
	mins := make([]int, n)
	m := math.MaxInt64
	for i := n - 1; i >= 0; i-- {
		mins[i] = m
		m = min(m, nums[i])
	}

	// Cumulative max from the left.
	maxs := make([]int, n)
	m = math.MinInt64
	for i := 0; i < n; i++ {
		maxs[i] = m
		m = max(m, nums[i])
	}

	// Find left index.
	l := 0
	for ; l+1 < n && nums[l] <= nums[l+1] && nums[l] <= mins[l]; l++ {
	}

	// Find right index.
	r := n - 1
	for ; r >= l && r-1 >= 0 && nums[r-1] <= nums[r] && nums[r] >= maxs[r]; r-- {
	}

	return r - l + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
