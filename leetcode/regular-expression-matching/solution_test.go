package solution

import (
	"strings"
	"testing"
)

func TestIsMatch(t *testing.T) {
	cases := []struct {
		String  string
		Pattern string
		Expect  bool
	}{
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b*", true},
		{"mississippi", "mis*is*p*.", false},

		{"a", "ab*a", false},

		{strings.Repeat("a", 500) + "b", "a*a*a*a*", false},
		{strings.Repeat("a", 500) + "b", "a*aa*aa*aa*ab", true},
	}
	for _, c := range cases {
		if got := isMatch(c.String, c.Pattern); got != c.Expect {
			t.Errorf("isMatch(%q, %q) = %v; expect %v", c.String, c.Pattern, got, c.Expect)
		}
	}
}
