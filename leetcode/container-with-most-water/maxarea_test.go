package maxarea

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect int
	}{
		{
			Input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Expect: 49,
		},
		{
			Input:  []int{0, 0},
			Expect: 0,
		},
	}
	for _, c := range cases {
		got := maxArea(c.Input)
		if got != c.Expect {
			t.Fatalf("maxArea(%v) = %d; expect %d", c.Input, got, c.Expect)
		}
	}
}

func TestMaxAreaRandom(t *testing.T) {
	for trial := 0; trial < 1000; trial++ {
		height := testcase()
		expect := naive(height)
		got := maxArea(height)
		if !reflect.DeepEqual(got, expect) {
			t.FailNow()
		}
	}
}

func naive(height []int) int {
	maxa := 0
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			a := abs(i-j) * min(height[i], height[j])
			if a > maxa {
				maxa = a
			}
		}
	}
	return maxa
}

func testcase() []int {
	n, limit := 100, 1000
	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i] = rand.Intn(limit)
	}
	return height
}
