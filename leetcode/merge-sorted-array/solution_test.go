package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		A, B []int
	}{
		{A: []int{1, 2, 3}, B: []int{2, 5, 6}},
		{A: []int{1, 2, 3}, B: []int{4, 5, 6}},
		{A: []int{1, 2, 3}, B: []int{}},
		{A: []int{}, B: []int{4, 5, 6}},
		{A: []int{1, 3, 5}, B: []int{2, 4, 6}},
		{A: []int{}, B: []int{}},
		{A: []int{3}, B: []int{2, 5, 6}},
	}
	for _, c := range cases {
		test(t, c.A, c.B)
		test(t, c.B, c.A)
	}
}

func test(t *testing.T, a, b []int) {
	m, n := len(a), len(b)

	// Compute with our implementation.
	nums1 := make([]int, m+n)
	nums2 := make([]int, n)
	copy(nums1, a)
	copy(nums2, b)
	merge(nums1, m, nums2, n)

	// Compute in a blatantly correct way.
	expect := append([]int{}, a...)
	expect = append(expect, b...)
	sort.Ints(expect)

	if !reflect.DeepEqual(nums1, expect) {
		t.Errorf("merge(%v, %d, %v, %d) got %v; expect %v", a, m, b, n, nums1, expect)
	}
}
