package solution

import (
	"math"
	"math/rand"
	"sort"
	"testing"
)

const epsilon = 0.00001

func TestMedianFinderExamples(t *testing.T) {
	cases := []struct {
		Nums   []int
		Expect float64
	}{
		{[]int{2, 3, 4}, 3},
		{[]int{2, 3}, 2.5},
	}
	for _, c := range cases {
		if got := StreamingMedian(c.Nums); got != c.Expect {
			t.Errorf("StreamingMedian(%v) = %v; expect %v", c.Nums, got, c.Expect)
		}
		if got := NaiveMedian(c.Nums); got != c.Expect {
			t.Errorf("NaiveMedian(%v) = %v; expect %v", c.Nums, got, c.Expect)
		}
	}
}

func TestMedianFinderRandom(t *testing.T) {
	for trial := 0; trial < 100; trial++ {
		f := Constructor()
		nums := []int{}
		for i := 0; i < 100; i++ {
			num := rand.Intn(100000)
			nums = append(nums, num)
			f.AddNum(num)

			got := f.FindMedian()
			expect := NaiveMedian(nums)
			if math.Abs(got-expect) > epsilon {
				t.Fatalf("nums=%v got=%v expect%v", nums, got, expect)
			}
		}
	}
}

func StreamingMedian(nums []int) float64 {
	f := Constructor()
	for _, num := range nums {
		f.AddNum(num)
	}
	return f.FindMedian()
}

func NaiveMedian(nums []int) float64 {
	n := len(nums)
	sorted := append([]int{}, nums...)
	sort.Ints(sorted)
	if n%2 == 0 {
		return float64(sorted[n/2-1]+sorted[n/2]) / 2
	}
	return float64(sorted[n/2])
}
