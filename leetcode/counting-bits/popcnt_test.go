package popcnt

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	cases := []struct {
		Input  int
		Expect []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
		{0, []int{0}},
	}
	for _, c := range cases {
		got := countBits(c.Input)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("countBits(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
