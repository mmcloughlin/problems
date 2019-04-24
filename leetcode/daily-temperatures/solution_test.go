package solution

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	cases := []struct {
		T      []int
		Expect []int
	}{
		{
			T:      []int{73, 74, 75, 71, 69, 72, 76, 73},
			Expect: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			T:      []int{70, 71, 72},
			Expect: []int{1, 1, 0},
		},
		{
			T:      []int{72, 71, 70},
			Expect: []int{0, 0, 0},
		},
		{
			T:      []int{72},
			Expect: []int{0},
		},
		{
			T:      []int{},
			Expect: []int{},
		},
	}
	for _, c := range cases {
		if got := dailyTemperatures(c.T); !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("dailyTemperatures(%v) = %v' expect %v", c.T, got, c.Expect)
		}
	}
}
