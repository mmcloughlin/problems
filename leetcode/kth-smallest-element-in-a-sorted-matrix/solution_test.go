package solution

import (
	"sort"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	matrix := [][]int{
		[]int{1, 5, 9},
		[]int{10, 11, 13},
		[]int{12, 13, 15},
	}
	sorted := sortedmatrix(matrix)
	for i, expect := range sorted {
		got := kthSmallest(matrix, i+1)
		if got != expect {
			t.Errorf("kthSmallest(..., %v) = %v; expect %v", i+1, got, expect)
		}
	}
}

func sortedmatrix(matrix [][]int) []int {
	concat := []int{}
	for _, row := range matrix {
		concat = append(concat, row...)
	}
	sort.Ints(concat)
	return concat
}
