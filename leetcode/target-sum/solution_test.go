package solution

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	cases := []struct {
		Input  []int
		S      int
		Expect int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
	}
	for _, c := range cases {
		if got := findTargetSumWays(c.Input, c.S); got != c.Expect {
			t.Errorf("findTargetSumWays(%v, %d) = %d; expect %d", c.Input, c.S, got, c.Expect)
		}
	}
}
