package search

var notfound = []int{-1, -1}

func searchRange(nums []int, target int) []int {
	n := len(nums)

	b := find(n, func(i int) bool {
		return nums[i] >= target
	})

	if b == n || nums[b] != target {
		return notfound
	}

	e := find(n-b, func(i int) bool {
		return nums[b+i] > target
	})

	return []int{b, b + e - 1}
}

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
