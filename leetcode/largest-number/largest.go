package largest

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	if sum(nums) == 0 {
		return "0"
	}

	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}

	sort.Slice(ss, func(i, j int) bool { return !less(ss[i], ss[j]) })

	return strings.Join(ss, "")
}

func less(a, b string) bool {
	m := min(len(a), len(b))
	for i := 0; i < m; i++ {
		if a[i] < b[i] {
			return true
		}
		if a[i] > b[i] {
			return false
		}
	}
	if len(a) < len(b) {
		return less(a, b[m:])
	}
	if len(a) > len(b) {
		return less(a[m:], b)
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
