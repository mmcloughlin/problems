package solution

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		Input  string
		Expect bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"abba", true},
		{"abbb", false},
		{"a....b::;   b??a", true},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"0P", false},
	}
	for _, c := range cases {
		if got := isPalindrome(c.Input); got != c.Expect {
			t.Errorf("isPalindrome(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
