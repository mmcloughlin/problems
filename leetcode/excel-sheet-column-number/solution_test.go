package solution

import "testing"

func TestTitleToNumber(t *testing.T) {
	cases := []struct {
		Input  string
		Expect int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}
	for _, c := range cases {
		if got := titleToNumber(c.Input); got != c.Expect {
			t.Errorf("titleToNumber(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
