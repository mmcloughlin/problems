package solution

import "testing"

func TestLongestSubstring(t *testing.T) {
	cases := []struct {
		Input  string
		K      int
		Expect int
	}{
		{"aaabb", 3, 3},
		{"ababbc", 2, 5},
		{"baaabcb", 3, 3},
	}
	for _, c := range cases {
		if got := longestSubstring(c.Input, c.K); got != c.Expect {
			t.Errorf("longestSubstring(%v, %v) = %v; expect %v", c.Input, c.K, got, c.Expect)
		}
	}
}
