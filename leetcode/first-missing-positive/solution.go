package solution

import "math"

func firstMissingPositive(nums []int) int {
	const missing = math.MinInt64
	n := len(nums)
	for i := 0; i < n; i++ {
		x := nums[i]
		j := x - 1 // index it belongs in
		switch {
		case x <= 0, x > n:
			nums[i] = missing
		case j < i:
			nums[j] = x
			nums[i] = missing
		case j > i:
			if nums[j] != x {
				nums[i], nums[j] = nums[j], x
				i--
			} else {
				nums[i] = missing
			}
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] == missing {
			return i + 1
		}
	}
	return n + 1
}
