package solution

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	cases := []struct {
		Input  string
		Expect int
	}{
		{"(()", 2},
		{")()())", 4},
		{"()(((((())(", 4},
		{"()", 2},
		{"(()())", 6},
		{"()(", 2},
		{")()", 2},
		{"(((()(()", 2},
		{"(()()", 4},
		{"(())()(()((", 6},
		{"()()))))))))()()()((()))()", 14},
		{"()()(()(((()))()))()))))()(())))()(()())()()()))())))())())))(()()()))))()((()(())(())))((()())(()()()((((()(())))))((()()((())(())(()(())))))()()())(())(()())((()())()(((())))()(()()))", 96},
	}
	for _, c := range cases {
		if got := longestValidParentheses(c.Input); got != c.Expect {
			t.Errorf("longestValidParentheses(%q) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
