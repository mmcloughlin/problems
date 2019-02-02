package rpn

import "testing"

func TestEvalRPN(t *testing.T) {
	cases := []struct {
		Tokens []string
		Expect int
	}{
		{
			Tokens: []string{"2", "1", "+", "3", "*"},
			Expect: 9,
		},
		{
			Tokens: []string{"4", "13", "5", "/", "+"},
			Expect: 6,
		},
		{
			Tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			Expect: 22,
		},
	}
	for _, c := range cases {
		got := evalRPN(c.Tokens)
		if got != c.Expect {
			t.Errorf("evalRPN(%v) = %v; expect %v", c.Tokens, got, c.Expect)
		}
	}
}
