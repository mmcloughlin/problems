package solution

import (
	"math/rand"
	"testing"
)

func TestFindDuplicateCases(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect int
	}{
		{
			Input:  []int{1, 3, 4, 2, 2},
			Expect: 2,
		},
		{
			Input:  []int{1, 2, 4, 2, 2},
			Expect: 2,
		},
		{
			Input:  []int{3, 1, 3, 4, 2},
			Expect: 3,
		},
	}
	for _, c := range cases {
		if got := findDuplicate(c.Input); got != c.Expect {
			t.Errorf("findDuplicate(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}

func TestFindDuplicateAll(t *testing.T) {
	n := 1000
	for dupe := 1; dupe <= n; dupe++ {
		nums := make([]int, n+1)
		nums[0] = dupe
		for i := 1; i <= n; i++ {
			nums[i] = i
		}

		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		for i := 0; i < 5; i++ {
			if got := findDuplicate(nums); got != dupe {
				t.Errorf("findDuplicate(%v) = %v; expect %v", nums, got, dupe)
			}

			for {
				j := rand.Intn(len(nums))
				if nums[j] != dupe {
					nums[j] = dupe
					break
				}
			}
		}
	}
}
