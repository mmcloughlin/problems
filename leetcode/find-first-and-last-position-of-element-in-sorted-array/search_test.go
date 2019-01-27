package search

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		Input  []int
		Target int
		Expect []int
	}{
		{
			Input:  []int{5, 7, 7, 8, 8, 10},
			Target: 6,
			Expect: notfound,
		},
		{
			Input:  []int{5, 7, 7, 8, 8, 10},
			Target: 8,
			Expect: []int{3, 4},
		},
		{
			Input:  []int{5, 7, 7, 8, 8, 10},
			Target: 7,
			Expect: []int{1, 2},
		},
		{
			Input:  []int{1, 2, 3, 4, 5, 6},
			Target: 5,
			Expect: []int{4, 4},
		},
		{
			Input:  []int{},
			Target: 5,
			Expect: notfound,
		},
		{
			Input:  []int{5},
			Target: 5,
			Expect: []int{0, 0},
		},
	}
	for _, c := range cases {
		got := searchRange(c.Input, c.Target)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("searchRange(%v, %v) = %v; expect %v", c.Input, c.Target, got, c.Expect)
		}
	}
}
