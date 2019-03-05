package solution

import "testing"

func TestIsMatch(t *testing.T) {
	cases := []struct {
		Input   string
		Pattern string
		Expect  bool
	}{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
		{
			"bbbababbbbabbbbababbaaabbaababbbaabbbaaaabbbaaaabb",
			"*b******bb*b*bbbbb*ba",
			false,
		},
		{
			"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
			"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb",
			false,
		},
	}
	for _, c := range cases {
		if got := isMatch(c.Input, c.Pattern); got != c.Expect {
			t.Errorf("isMatch(%v, %v) = %v; expect %v", c.Input, c.Pattern, got, c.Expect)
		}
	}
}
