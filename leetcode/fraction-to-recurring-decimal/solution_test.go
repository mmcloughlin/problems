package solution

import "testing"

func TestFractionToDecimal(t *testing.T) {
	cases := []struct {
		Numerator   int
		Denominator int
		Expect      string
	}{
		{1, 2, "0.5"},
		{2, 1, "2"},
		{2, 3, "0.(6)"},
		{811, 70, "11.5(857142)"},
		{9, 3, "3"},
		{9, 3, "3"},
		{-1, 7, "-0.(142857)"},
		{1, -7, "-0.(142857)"},
		{-1, -7, "0.(142857)"},
		{0, -5, "0"},
	}
	for _, c := range cases {
		if got := fractionToDecimal(c.Numerator, c.Denominator); got != c.Expect {
			t.Errorf("%v / %v = %v; expect %v", c.Numerator, c.Denominator, got, c.Expect)
		}
	}
}
