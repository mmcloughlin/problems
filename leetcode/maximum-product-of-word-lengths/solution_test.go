package solution

import "testing"

func TestMaxProduct(t *testing.T) {
	cases := []struct {
		Words  []string
		Expect int
	}{
		{
			Words:  []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
			Expect: 16,
		},
		{
			Words:  []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
			Expect: 4,
		},
		{
			Words:  []string{"a", "aa", "aaa", "aaaa"},
			Expect: 0,
		},
		{
			Words:  []string{"eae", "ea", "aaf", "bda", "fcf", "dc", "ac", "ce", "cefde", "dabae"},
			Expect: 15,
		},
	}
	for _, c := range cases {
		if got := maxProduct(c.Words); got != c.Expect {
			t.Errorf("maxProduct(%v) = %v; expect %v", c.Words, got, c.Expect)
		}
	}
}
