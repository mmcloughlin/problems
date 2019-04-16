package solution

import (
	"math/rand"
	"testing"
)

func TestFirstMissingPositiveExamples(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{1, 1}, 2},
		{[]int{2, 2}, 1},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}
	for _, c := range cases {
		if got := firstMissingPositive(c.Input); got != c.Expect {
			t.Fatalf("firstMissingPositive(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}

func TestFirstMissingPositiveRandom(t *testing.T) {
	for trial := 0; trial < 100; trial++ {
		n := 3 + rand.Intn(128)
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			nums[i] = i + 1
		}

		m := 1 + rand.Intn(n)
		nums[m-1] = 2 * n

		rand.Shuffle(n, func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		input := append([]int{}, nums...)
		got := firstMissingPositive(input)
		if m != got {
			t.Fatalf("firstMissingPositive(%v) = %d; expect %d", nums, got, m)
		}
	}
}
