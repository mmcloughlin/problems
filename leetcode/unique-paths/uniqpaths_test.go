package uniqpaths

import "testing"

func TestUniquePaths(t *testing.T) {
	cases := []struct {
		M, N   int
		Expect int
	}{
		{3, 2, 3},
		{7, 3, 28},
		{3, 7, 28},
		{1, 1, 1},
		{1, 100, 1},
		{13, 23, 548354040},
	}
	for _, c := range cases {
		got := uniquePaths(c.M, c.N)
		if got != c.Expect {
			t.Errorf("uniquePaths(%v, %v) = %v; expect %v", c.M, c.N, got, c.Expect)
		}
	}
}
