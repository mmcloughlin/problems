package maxprod

import "testing"

func TestMaxProduct(t *testing.T) {
	cases := []struct {
		Nums   []int
		Expect int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-10}, -10},
		{[]int{-10, 100}, 100},
		{[]int{-10, 100, -10}, 10000},
	}
	for _, c := range cases {
		got := maxProduct(c.Nums)
		if got != c.Expect {
			t.Errorf("maxProduct(%v) = %v; expect %v", c.Nums, got, c.Expect)
		}
	}
}
