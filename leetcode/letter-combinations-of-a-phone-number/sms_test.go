package sms

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	cases := []struct {
		Input  string
		Expect []string
	}{
		{
			Input:  "23",
			Expect: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			Input:  "0",
			Expect: []string{},
		},
		{
			Input:  "hello",
			Expect: []string{},
		},
		{
			Input:  "",
			Expect: []string{},
		},
	}
	for _, c := range cases {
		got := letterCombinations(c.Input)
		sort.Strings(got)
		sort.Strings(c.Expect)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Logf("got\t%v", got)
			t.Logf("expect\t%v", c.Expect)
			t.Fail()
		}
	}
}
