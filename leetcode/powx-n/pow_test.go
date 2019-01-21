package pow

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	cases := []struct {
		X      float64
		N      int
		Expect float64
	}{
		{2, 10, 1024},
		{2.1, 3, 9.261},
		{2, -2, 0.25},
		{10, 0, 1},
	}
	const epsilon = 1e-8
	for _, c := range cases {
		got := myPow(c.X, c.N)
		if math.Abs(got-c.Expect) > epsilon {
			t.Errorf("myPow(%v, %v) = %v; expect %v", c.X, c.N, got, c.Expect)
		}
	}
}
