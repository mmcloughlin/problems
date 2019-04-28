package solution

import "testing"

func TestIntegerBreak(t *testing.T) {
	cases := []struct {
		N      int
		Expect int
	}{
		{2, 1},
		{3, 2},
		{10, 36},
	}
	for _, c := range cases {
		if got := integerBreak(c.N); got != c.Expect {
			t.Errorf("integerBreak(%v) = %v; expect %v", c.N, got, c.Expect)
		}
	}
}
