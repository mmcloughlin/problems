package parens

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		Input  string
		Expect bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"[", false},
	}
	for _, c := range cases {
		got := isValid(c.Input)
		if got != c.Expect {
			t.Errorf("isValid(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
