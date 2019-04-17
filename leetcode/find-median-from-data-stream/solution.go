package solution

import "container/heap"

type MedianFinder struct {
	l, r *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		l: &MinHeap{},
		r: &MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// Add num to one side.
	if this.r.Len() > 0 && num >= this.r.Min() {
		heap.Push(this.r, num)
	} else {
		heap.Push(this.l, -num)
	}

	// Rebalance.
	for this.l.Len() < this.r.Len() {
		x := heap.Pop(this.r)
		heap.Push(this.l, -x.(int))
	}

	for this.l.Len() > this.r.Len()+1 {
		x := heap.Pop(this.l)
		heap.Push(this.r, -x.(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	switch {
	case this.l.Len() == this.r.Len()+1:
		return -float64(this.l.Min())
	case this.l.Len() == this.r.Len():
		return 0.5 * (-float64(this.l.Min()) + float64(this.r.Min()))
	default:
		panic("invariant violated")
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// MinHeap satisfies heap.Interface.
type MinHeap struct {
	nums []int
}

func (h MinHeap) Min() int {
	return h.nums[0]
}

func (h MinHeap) Len() int { return len(h.nums) }

func (h MinHeap) Less(i, j int) bool { return h.nums[i] < h.nums[j] }

func (h MinHeap) Swap(i, j int) { h.nums[i], h.nums[j] = h.nums[j], h.nums[i] }

func (h *MinHeap) Push(x interface{}) {
	h.nums = append(h.nums, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	t := len(h.nums) - 1
	x := h.nums[t]
	h.nums = h.nums[:t]
	return x
}
