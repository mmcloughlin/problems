package solution

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	cases := []struct {
		Input  []int
		K      int
		Expect []int
	}{
		{
			Input:  []int{1, 3, -1, -3, 5, 3, 6, 7},
			K:      3,
			Expect: []int{3, 3, 5, 5, 6, 7},
		},
	}
	for _, c := range cases {
		got := maxSlidingWindow(c.Input, c.K)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("maxSlidingWindow(%v, %v) = %v; expect %v", c.Input, c.K, got, c.Expect)
		}
	}
}

func TestMaxSlidingWindowFuzz(t *testing.T) {
	for trial := 0; trial < 1000; trial++ {
		nums := randslice(100)
		k := 1 + rand.Intn(50)
		expect := naive(nums, k)
		got := maxSlidingWindow(nums, k)
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("maxSlidingWindow(%v, %v) = %v; expect %v", nums, k, got, expect)
		}
	}
}

func randslice(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100000)
	}
	return nums
}

func naive(nums []int, k int) []int {
	m := []int{}
	for i := 0; i+k <= len(nums); i++ {
		m = append(m, slicemax(nums[i:i+k]))
	}
	return m
}

func slicemax(nums []int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}
