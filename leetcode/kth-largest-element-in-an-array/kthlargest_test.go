package kthlargest

import (
	"sort"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	cases := [][]int{
		{3, 2, 1, 5, 6, 4},
	}
	for _, c := range cases {
		sorted := append([]int{}, c...)
		sort.Ints(sorted)
		for k := 1; len(c)-k >= 0; k++ {
			got := findKthLargest(c, k)
			expect := sorted[len(c)-k]
			if got != expect {
				t.Errorf("findKthLargest(%v, %v) = %v; expect %v", c, k, got, expect)
			}
		}
	}
}
