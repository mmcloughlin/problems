package solution

func findDuplicate(nums []int) int {
	n := len(nums) - 1
	l, r := 1, n
	for l < r {
		m := l + (r-l)/2

		c := 0
		for _, num := range nums {
			if l <= num && num <= m {
				c++
			}
		}

		if c <= m-l+1 {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
