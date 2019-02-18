package stock

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		Prices []int
		Expect int
	}{
		{
			Prices: []int{7, 1, 5, 3, 6, 4},
			Expect: 5,
		},
		{
			Prices: []int{7, 6, 4, 3, 1},
			Expect: 0,
		},
		{
			Prices: []int{},
			Expect: 0,
		},
	}
	for _, c := range cases {
		if got := maxProfit(c.Prices); got != c.Expect {
			t.Errorf("maxProfit(%v) = %v; expect %v", c.Prices, got, c.Expect)
		}
	}
}
