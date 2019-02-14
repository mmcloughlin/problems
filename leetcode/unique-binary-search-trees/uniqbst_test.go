package uniqbst

import "testing"

func TestNumTrees(t *testing.T) {
	cases := []struct {
		N      int
		Expect int
	}{
		{3, 5},
		{0, 1},
		{1, 1},
	}
	for _, c := range cases {
		got := numTrees(c.N)
		if got != c.Expect {
			t.Errorf("numTrees(%v) = %v; expect %v", c.N, got, c.Expect)
		}
	}
}
