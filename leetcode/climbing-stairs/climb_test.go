package climb

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		N      int
		Expect int
	}{
		{2, 2},
		{3, 3},
	}
	for _, c := range cases {
		if got := climbStairs(c.N); got != c.Expect {
			t.Errorf("climbStairs(%v) = %v; expect %v", c.N, got, c.Expect)
		}
	}
}
