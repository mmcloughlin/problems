package solution

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		Prices []int
		Expect int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		{[]int{1, 2, 3, 4, 50}, 49},
		{[]int{5, 3, 1}, 0},
		{[]int{101, 102, 103, 1, 3, 4}, 4},
	}
	for _, c := range cases {
		if got := maxProfit(c.Prices); got != c.Expect {
			t.Errorf("maxProfit(%v) = %v; expect %v", c.Prices, got, c.Expect)
		}
	}
}
