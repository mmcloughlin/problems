package kthlargest

import "container/heap"

type MaxHeap struct {
	Size int
	Nums []int
}

func (h *MaxHeap) Insert(x int) {
	if h.Len() < h.Size {
		heap.Push(h, x)
	} else if h.Nums[0] < x {
		h.Nums[0] = x
		heap.Fix(h, 0)
	}
}

func (h *MaxHeap) Len() int { return len(h.Nums) }

func (h *MaxHeap) Less(i, j int) bool {
	return h.Nums[i] < h.Nums[j]
}

func (h *MaxHeap) Swap(i, j int) {
	h.Nums[i], h.Nums[j] = h.Nums[j], h.Nums[i]
}

func (h *MaxHeap) Push(x interface{}) {
	h.Nums = append(h.Nums, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	n := len(h.Nums)
	x := h.Nums[n-1]
	h.Nums = h.Nums[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MaxHeap{Size: k}
	for _, num := range nums {
		h.Insert(num)
	}
	return h.Nums[0]
}
