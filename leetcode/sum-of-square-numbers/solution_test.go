package solution

import (
	"math"
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	cases := []struct {
		C      int
		Expect bool
	}{
		{1*1 + 2*2, true},
		{3, false},
		{16, true},
		{math.MaxInt16*math.MaxInt16 + 3*3, true},
	}
	for _, c := range cases {
		if got := judgeSquareSum(c.C); got != c.Expect {
			t.Errorf("judgeSquareSum(%v) = %v; expect %v", c.C, got, c.Expect)
		}
	}
}
