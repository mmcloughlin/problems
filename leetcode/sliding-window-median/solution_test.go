package solution

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestMedianSlidingWindow(t *testing.T) {
	cases := []struct {
		Nums   []int
		K      int
		Expect []float64
	}{
		{
			Nums:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			K:      3,
			Expect: []float64{1, -1, -1, 3, 5, 6},
		},
		{
			Nums:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			K:      1,
			Expect: []float64{1, 3, -1, -3, 5, 3, 6, 7},
		},
		{
			Nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			K:      4,
			Expect: []float64{2.5, 3.5, 4.5, 5.5, 6.5, 7.5},
		},
	}
	for _, c := range cases {
		got := medianSlidingWindow(c.Nums, c.K)
		if !reflect.DeepEqual(c.Expect, got) {
			t.Errorf("medianSlidingWindow(%v, %v) = %v; expect %v", c.Nums, c.K, got, c.Expect)
		}
	}
}

func TestMedianSlidingWindowRandom(t *testing.T) {
	n, max := 100, 1000
	for trial := 0; trial < 100; trial++ {
		nums := randints(n, max)
		k := 1 + rand.Intn(n-1)
		got := medianSlidingWindow(nums, k)
		expect := naive(nums, k)
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("medianSlidingWindow(%v, %v) = %v; expect %v", nums, k, got, expect)
		}
	}
}

func randints(n, max int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(max)
	}
	return nums
}

func naive(nums []int, k int) []float64 {
	medians := []float64{}
	for i := 0; i+k <= len(nums); i++ {
		medians = append(medians, median(nums[i:i+k]))
	}
	return medians
}

func median(nums []int) float64 {
	sorted := append([]int{}, nums...)
	sort.Ints(sorted)
	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	}
	return (float64(sorted[n/2]) + float64(sorted[n/2-1])) / 2
}
