package div

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	cases := []struct {
		Dividend int
		Divisor  int
		Expect   int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{16, 2, 8},
		{64, 16, 4},
		{123, 1, 123},
		{0, 123, 0},
	}
	for _, c := range cases {
		got := divide(c.Dividend, c.Divisor)
		if got != c.Expect {
			t.Errorf("divide(%d, %d) = %d; expect %d", c.Dividend, c.Divisor, got, c.Expect)
		}
	}
}

func TestDivideOverflow(t *testing.T) {
	got := divide(1<<35, 2)
	if got != math.MaxInt32 {
		t.Fail()
	}
}
