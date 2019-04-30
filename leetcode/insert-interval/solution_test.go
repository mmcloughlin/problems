package solution

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		Intervals [][]int
		New       []int
		Expect    [][]int
	}{
		{
			Intervals: [][]int{{1, 3}, {6, 9}},
			New:       []int{2, 5},
			Expect:    [][]int{{1, 5}, {6, 9}},
		},
		{
			Intervals: [][]int{{1, 3}, {6, 9}},
			New:       []int{5, 8},
			Expect:    [][]int{{1, 3}, {5, 9}},
		},
		{
			Intervals: [][]int{{3, 5}, {12, 15}},
			New:       []int{7, 8},
			Expect:    [][]int{{3, 5}, {7, 8}, {12, 15}},
		},
		{
			Intervals: [][]int{{3, 5}, {12, 15}},
			New:       []int{6, 6},
			Expect:    [][]int{{3, 5}, {6, 6}, {12, 15}},
		},
	}
	for _, c := range cases {
		if got := insert(c.Intervals, c.New); !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("insert(%v, %v) = %v; expect %v", c.Intervals, c.New, got, c.Expect)
		}
	}
}
