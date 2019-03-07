package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestWordBreak(t *testing.T) {
	cases := []struct {
		Input  string
		Dict   []string
		Expect []string
	}{
		{
			Input: "catsanddog",
			Dict:  []string{"cat", "cats", "and", "sand", "dog"},
			Expect: []string{
				"cats and dog",
				"cat sand dog",
			},
		},
		{
			Input: "pineapplepenapple",
			Dict:  []string{"apple", "pen", "applepen", "pine", "pineapple"},
			Expect: []string{
				"pine apple pen apple",
				"pineapple pen apple",
				"pine applepen apple",
			},
		},
		{
			Input:  "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Dict:   []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			Expect: []string{},
		},
	}
	for _, c := range cases {
		got := wordBreak(c.Input, c.Dict)
		sort.Strings(got)
		sort.Strings(c.Expect)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("wordBreak(%v, ...) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
