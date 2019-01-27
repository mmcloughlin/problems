package search

// find returns the first integer at which predicate is true. It is assumed
// that predicate will be false for a run of values, and then true for all
// following. Returns n if predicate is always false.
func find(n int, predicate func(int) bool) int {
	// empty case
	if n == 0 {
		return 0
	}

	// if the last position is false, everything is false.
	if !predicate(n - 1) {
		return n
	}

	if predicate(0) {
		return 0
	}

	// invariant: predicate(l) is false and predicate(r) is true
	l, r := 0, n-1
	for r-l > 1 {
		m := l + (r-l)/2
		if predicate(m) {
			r = m
		} else {
			l = m
		}
	}

	return r
}

func search(nums []int, target int) int {
	n := len(nums)

	// find the position at which the array cycles around
	c := find(n, func(i int) bool {
		return nums[i] < nums[0]
	})

	// the sequence nums[i+c] is sorted
	i := find(n, func(i int) bool {
		return nums[(i+c)%n] >= target
	})

	if i == n {
		return -1
	}

	i = (i + c) % n
	if nums[i] == target {
		return i
	}

	return -1
}
