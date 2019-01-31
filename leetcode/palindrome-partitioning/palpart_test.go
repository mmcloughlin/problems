package palpart

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		Input  string
		Expect [][]string
	}{
		{
			Input: "aab",
			Expect: [][]string{
				{"aa", "b"},
				{"a", "a", "b"},
			},
		},
		{
			Input: "aaba",
			Expect: [][]string{
				{"a", "aba"},
				{"aa", "b", "a"},
				{"a", "a", "b", "a"},
			},
		},
		{
			Input: "a",
			Expect: [][]string{
				{"a"},
			},
		},
	}
	for _, c := range cases {
		got := partition(c.Input)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("partition(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
