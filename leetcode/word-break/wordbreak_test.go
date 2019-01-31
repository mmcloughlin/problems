package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	cases := []struct {
		Input  string
		Dict   []string
		Expect bool
	}{
		{
			Input:  "leetcode",
			Dict:   []string{"leet", "code"},
			Expect: true,
		},
		{
			Input:  "applepenapple",
			Dict:   []string{"apple", "pen"},
			Expect: true,
		},
		{
			Input:  "catsandog",
			Dict:   []string{"cats", "dog", "sand", "and", "cat"},
			Expect: false,
		},
		{
			Input:  "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			Dict:   []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			Expect: false,
		},
	}
	for _, c := range cases {
		got := wordBreak(c.Input, c.Dict)
		if got != c.Expect {
			t.Errorf("wordBreak(%v, %v) = %v; expect %v", c.Input, c.Dict, got, c.Expect)
		}
	}
}
