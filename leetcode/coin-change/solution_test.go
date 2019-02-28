package solution

import "testing"

func TestCoinChange(t *testing.T) {
	cases := []struct {
		Coins  []int
		Amount int
		Expect int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
	}
	for _, c := range cases {
		if got := coinChange(c.Coins, c.Amount); got != c.Expect {
			t.Errorf("coinChange(%v, %v) = %v; expect %v", c.Coins, c.Amount, got, c.Expect)
		}
	}
}
