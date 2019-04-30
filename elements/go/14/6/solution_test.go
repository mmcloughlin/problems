package solution

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	cases := []struct {
		Input  []Interval
		Add    Interval
		Expect []Interval
	}{
		{
			// Book example.
			Input: []Interval{
				{-4, -1},
				{0, 2},
				{3, 6},
				{7, 9},
				{11, 12},
				{14, 17},
			},
			Add: Interval{1, 8},
			Expect: []Interval{
				{-4, -1},
				{0, 9},
				{11, 12},
				{14, 17},
			},
		},
		{
			// Empty.
			Input:  []Interval{},
			Add:    Interval{1, 8},
			Expect: []Interval{{1, 8}},
		},
		{
			// Pre.
			Input: []Interval{
				{10, 12},
				{14, 19},
			},
			Add: Interval{1, 8},
			Expect: []Interval{
				{1, 8},
				{10, 12},
				{14, 19},
			},
		},
		{
			// Post.
			Input: []Interval{
				{1, 8},
				{10, 12},
			},
			Add: Interval{14, 19},
			Expect: []Interval{
				{1, 8},
				{10, 12},
				{14, 19},
			},
		},
		{
			// Total overlap.
			Input: []Interval{
				{10, 12},
				{14, 19},
			},
			Add: Interval{5, 30},
			Expect: []Interval{
				{5, 30},
			},
		},
		{
			// Partial overlap.
			Input: []Interval{
				{10, 12},
				{14, 19},
			},
			Add: Interval{11, 30},
			Expect: []Interval{
				{10, 30},
			},
		},
	}
	for _, c := range cases {
		if got := Union(c.Input, c.Add); !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("Union(%v, %v) = %v; expect %v", c.Input, c.Add, got, c.Expect)
		}
	}
}
