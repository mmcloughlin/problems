package solution

import "testing"

func TestCountSubstrings(t *testing.T) {
	cases := []struct {
		Input  string
		Expect int
	}{
		{"abc", 3},
		{"a", 1},
		{"aaa", 6},
		{"aba", 4},
		{"baa", 4},
	}
	for _, c := range cases {
		if got := countSubstrings(c.Input); got != c.Expect {
			t.Errorf("countSubstrings(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
