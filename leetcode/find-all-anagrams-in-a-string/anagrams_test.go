package anagrams

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	cases := []struct {
		S, P   string
		Expect []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}
	for _, c := range cases {
		if got := findAnagrams(c.S, c.P); !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("findAnagrams(%q, %q) = %v; expect %v", c.S, c.P, got, c.Expect)
		}
	}
}
