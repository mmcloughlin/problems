package solution

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		Prices []int
		Expect int
	}{
		{
			Prices: []int{7, 1, 5, 3, 6, 4},
			Expect: 7,
		},
		{
			Prices: []int{7, 6, 5, 4, 3, 2},
			Expect: 0,
		},
		{
			Prices: []int{1, 2, 3, 4, 5},
			Expect: 4,
		},
	}
	for _, c := range cases {
		if got := maxProfit(c.Prices); got != c.Expect {
			t.Fatalf("maxProfit(%v) = %v; expect %v", c.Prices, got, c.Expect)
		}
	}
}
