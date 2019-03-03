package solution

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestCountSmallerExample(t *testing.T) {
	input := []int{5, 2, 6, 1}
	expect := []int{2, 1, 1, 0}
	got := countSmaller(input)
	assertequal(t, got, expect)
}

func TestCountSmallerFuzz(t *testing.T) {
	n := 1000
	nums := make([]int, n)
	for trial := 0; trial < 1000; trial++ {
		for i := 0; i < n; i++ {
			nums[i] = int(rand.Int31())
		}
		got := countSmaller(nums)
		expect := naive(nums)
		assertequal(t, got, expect)
	}
}

// naive implementation of countSmaller.
func naive(nums []int) []int {
	n := len(nums)
	counts := make([]int, n)
	for i, num := range nums {
		for j := i + 1; j < n; j++ {
			if nums[j] < num {
				counts[i]++
			}
		}
	}
	return counts
}

func assertequal(t *testing.T, got, expect []int) {
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("got %v; expect %v", got, expect)
	}
}
