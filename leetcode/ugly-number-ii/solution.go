package solution

import "container/heap"

func nthUglyNumber(n int) int {
	// Initialize a heap.
	h := &minheap{
		nums: nil,
		set:  map[int]bool{},
	}
	heap.Push(h, 1)

	// Pull items off it.
	var ugly int
	for i := 0; i < n; i++ {
		ugly = heap.Pop(h).(int)
		for _, p := range []int{2, 3, 5} {
			nxt := ugly * p
			if !h.set[nxt] {
				heap.Push(h, nxt)
			}
		}
	}

	return ugly
}

type minheap struct {
	nums []int
	set  map[int]bool
}

func (h minheap) Len() int { return len(h.nums) }

func (h minheap) Less(i, j int) bool { return h.nums[i] < h.nums[j] }

func (h *minheap) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *minheap) Push(x interface{}) {
	n := x.(int)
	h.nums = append(h.nums, n)
	h.set[n] = true
}

func (h *minheap) Pop() interface{} {
	top := len(h.nums) - 1
	n := h.nums[top]
	h.nums = h.nums[:top]
	delete(h.set, n)
	return n
}
