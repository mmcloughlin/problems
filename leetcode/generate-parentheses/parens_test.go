package parens

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	cases := []struct {
		N      int
		Expect []string
	}{
		{N: 0, Expect: []string{""}},
		{
			N: 3,
			Expect: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
	}
	for _, c := range cases {
		got := generateParenthesis(c.N)
		sort.Strings(got)
		sort.Strings(c.Expect)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Logf("got\n%#v", got)
			t.Logf("expect\n%#v", c.Expect)
			t.Fail()
		}
	}
}
