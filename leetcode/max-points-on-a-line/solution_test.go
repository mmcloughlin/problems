package solution

import (
	"math/rand"
	"testing"
)

func TestMaxPointsExamples(t *testing.T) {
	cases := []struct {
		Input  [][]int
		Expect int
	}{
		// Examples from problem statement.
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		{[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},

		{[][]int{{84, 250}, {0, 0}, {1, 0}, {0, -70}, {0, -70}, {1, -1}, {21, 10}, {42, 90}, {-42, -230}}, 6},

		// Horizontal.
		{[][]int{{1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}}, 5},

		// Vertical.
		{[][]int{{3, 1}, {3, 2}, {3, 3}, {3, 4}, {3, 5}, {10, 3}}, 5},

		// Edge cases.
		{[][]int{}, 0},
		{[][]int{{0, 0}, {0, 0}}, 2},
		{[][]int{{1, 1}, {1, 1}, {2, 3}}, 3},
	}

	n := len(cases)
	for i := 0; i < n; i++ {
		c := cases[i]
		duped := [][]int{}
		for _, p := range c.Input {
			duped = append(duped, p)
			duped = append(duped, p)
		}
		c.Input = duped
		rand.Shuffle(len(c.Input), func(i, j int) {
			c.Input[i], c.Input[j] = c.Input[j], c.Input[i]
		})
	}

	for _, c := range cases {
		if got := maxPoints(c.Input); got != c.Expect {
			t.Fatalf("maxPoints(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
