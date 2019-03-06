package solution

import (
	"math/rand"
	"sort"
	"testing"
)

func TestFindMedianSortedArraysCases(t *testing.T) {
	assertcase(t, []int{1}, []int{2}, 1.5)
	assertcase(t, []int{1, 3}, []int{2}, 2.0)
	assertcase(t, []int{1, 2}, []int{3, 4}, 2.5)
	assertcase(t, []int{1}, []int{2, 3, 4}, 2.5)
	assertcase(t, []int{1, 4}, []int{2, 3}, 2.5)
}

func TestFindMedianSortedArraysFuzz(t *testing.T) {
	for trial := 0; trial < 1000; trial++ {
		x, y, expect := generate()
		assertcase(t, x, y, expect)
	}
}

func assertcase(t *testing.T, x, y []int, expect float64) {
	for i := 0; i < 2; i++ {
		if got := findMedianSortedArrays(x, y); got != expect {
			t.Fatalf("findMedianSortedArrays(%v, %v) = %v; expect %v", x, y, got, expect)
		}
		x, y = y, x
	}
}

func generate() ([]int, []int, float64) {
	// generate the merged array
	t := 2 + rand.Intn(1000)
	merged := make([]int, t)
	for i := 0; i < t; i++ {
		merged[i] = rand.Intn(10000)
	}

	// split it into two
	m := 1 + rand.Intn(t-1)
	a := append([]int{}, merged[:m]...)
	b := append([]int{}, merged[m:]...)

	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(merged)

	// compute the median of the entire thing
	return a, b, median(merged)
}

func median(nums []int) float64 {
	mid := len(nums) / 2
	if len(nums)%2 == 1 {
		return float64(nums[mid])
	}
	return float64(nums[mid-1]+nums[mid]) / 2
}
