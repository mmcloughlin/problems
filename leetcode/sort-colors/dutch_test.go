package dutch

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortColors(t *testing.T) {
	cases := [][]int{
		{2, 0, 2, 1, 1, 0},
		{0},
		{1},
		{2},
		{0, 0, 0},
		{0, 1, 2},
		{1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2},
	}
	for _, colors := range cases {
		expect := append([]int{}, colors...)
		sort.Ints(expect)
		sortColors(colors)
		t.Logf("got %v", colors)
		if !reflect.DeepEqual(expect, colors) {
			t.Errorf("expect %v", expect)
		}
	}
}
