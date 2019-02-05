package peak

import (
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 1, 1},
		{1, 2, 3, 1, 1, 1, 1},
		{1, 2, 1, 3, 5, 6, 4},
		{1},
		{1, 2},
		{63, -82, -27, -40, -57, -92, 44, -90, -26, -45, -23, 89, 55, 56, 4, 17, 24, -53, -44, 80, 26, -77, -5, -34, 99, -66, -9, 5, -32, 36, -65, -86, 71, 28, -19, 73, -52, 20, 43, 27, 23, 3, 16, 85, 40, -48, 93, -43, -17, 0},
	}
	for _, c := range cases {
		i := findPeakElement(c)
		if !ispeak(c, i) {
			t.Logf("%d %d %d", lookup(c, i-1), lookup(c, i), lookup(c, i+1))
			t.Errorf("findPeakElement(%v) = %d", c, i)
		}
	}
}
