package minwin

import "testing"

func TestMinWindow(t *testing.T) {
	cases := []struct {
		S      string
		T      string
		Expect string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"ABC", "ABC", "ABC"},
		{"ABC", "XYZ", ""},
		{"AABCBCBCBCBC", "AABBC", "AABCB"},
		{"AAABBBCCC", "ABC", "ABBBC"},
	}
	for _, c := range cases {
		got := minWindow(c.S, c.T)
		if got != c.Expect {
			t.Errorf("minWindow(%q, %q) = %q; expect %q", c.S, c.T, got, c.Expect)
		}
	}
}
