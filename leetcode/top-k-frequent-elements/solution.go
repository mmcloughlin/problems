package solution

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	// Compute counts.
	counts := map[int]int{}
	for _, num := range nums {
		counts[num]++
	}

	// Add into heap.
	t := topk{
		k:     k,
		count: counts,
	}

	for num, count := range counts {
		t.Add(num, count)
	}

	result := append([]int{}, t.nums...)
	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}
	return result
}

type topk struct {
	k     int
	nums  []int
	count map[int]int
}

func (t *topk) Add(n, c int) {
	t.count[n] = c
	if t.Len() < t.k {
		heap.Push(t, n)
	} else if t.count[t.nums[0]] < c {
		t.nums[0] = n
		heap.Fix(t, 0)
	}
}

func (t *topk) Len() int           { return len(t.nums) }
func (t *topk) Less(i, j int) bool { return t.count[t.nums[i]] < t.count[t.nums[j]] }
func (t *topk) Swap(i, j int)      { t.nums[i], t.nums[j] = t.nums[j], t.nums[i] }

func (t *topk) Push(x interface{}) {
	t.nums = append(t.nums, x.(int))
}

func (t *topk) Pop() interface{} {
	top := t.Len()
	x := t.nums[top]
	t.nums = t.nums[:top]
	return x
}
