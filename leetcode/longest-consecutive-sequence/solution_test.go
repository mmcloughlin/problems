package solution

import (
	"math/rand"
	"testing"
)

func TestLongestConsecutiveExample(t *testing.T) {
	input := []int{100, 4, 200, 1, 3, 2}
	expect := 4
	got := longestConsecutive(input)
	if longestConsecutive(input) != expect {
		t.Fatalf("got %v; expected %v", got, expect)
	}
}

func TestLongestConsecutive(t *testing.T) {
	for trial := 0; trial < 1000; trial++ {
		nums, expect := generate()
		got := longestConsecutive(nums)
		if got != expect {
			t.Fatalf("longestConsecutive(%v) = %v; expect %v", nums, got, expect)
		}
	}
}

func generate() ([]int, int) {
	N := 10
	nums := []int{}
	max := 1
	base := 0
	for i := 0; i < N; i++ {
		base += 1 + rand.Intn(1000)
		l := 1 + rand.Intn(255)
		for j := 0; j < l; j++ {
			nums = append(nums, base)
			base++
		}
		if l > max {
			max = l
		}
	}
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums, max
}
