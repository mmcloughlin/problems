package perfectsquares

import "testing"

func TestNumSquares(t *testing.T) {
	cases := []struct {
		N      int
		Expect int
	}{
		{12, 3},
		{13, 2},
		{9, 1},
	}
	for _, c := range cases {
		got := numSquares(c.N)
		if got != c.Expect {
			t.Errorf("numSquares(%v) = %v; expect %v", c.N, got, c.Expect)
		}
	}
}
