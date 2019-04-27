package solution

import "container/heap"

func medianSlidingWindow(nums []int, k int) []float64 {
	// Initialize median stream.
	s := NewMedianStream(nums)
	for i := 0; i < k; i++ {
		s.Push(i)
	}

	// Produce median values.
	m := []float64{s.Median()}
	for i := k; i < len(nums); i++ {
		s.Pop(i - k)
		s.Push(i)
		m = append(m, s.Median())
	}

	return m
}

// MedianStream maintains a streaming median.
type MedianStream struct {
	nums   []int
	lo, hi minheap
}

func NewMedianStream(nums []int) MedianStream {
	return MedianStream{
		nums: nums,
		lo:   newminheap(nums, -1),
		hi:   newminheap(nums, 1),
	}
}

func (m *MedianStream) Push(i int) {
	if m.Len() == 0 {
		m.lo.Add(i)
		return
	}

	x := m.nums[i]
	if x <= m.lo.Top() {
		m.lo.Add(i)
	} else {
		m.hi.Add(i)
	}

	m.fixup()
}

func (m *MedianStream) Pop(i int) {
	m.lo.Delete(i)
	m.hi.Delete(i)
	m.fixup()
}

func (m MedianStream) Median() float64 {
	switch m.lo.Len() - m.hi.Len() {
	case 0:
		return (float64(m.lo.Top()) + float64(m.hi.Top())) / 2.0
	case 1:
		return float64(m.lo.Top())
	}
	panic("invariant violated")
}

func (m *MedianStream) fixup() {
	for m.lo.Len() > m.hi.Len()+1 {
		m.hi.Add(m.lo.PopTop())
	}
	for m.lo.Len() < m.hi.Len() {
		m.lo.Add(m.hi.PopTop())
	}
}

func (m *MedianStream) Len() int {
	return m.lo.Len() + m.hi.Len()
}

// minheap is a min heap if indexes into an underlying array.
type minheap struct {
	sgn  int
	nums []int       // underlying array
	idx  []int       // indexes into nums in heap order
	pos  map[int]int // mapping from index to position in idx array
}

func newminheap(nums []int, sgn int) minheap {
	return minheap{
		sgn:  sgn,
		nums: nums,
		idx:  nil,
		pos:  map[int]int{},
	}
}

func (h minheap) Top() int {
	return h.nums[h.idx[0]]
}

func (h *minheap) PopTop() int {
	return heap.Pop(h).(int)
}

func (h *minheap) Add(i int) {
	heap.Push(h, i)
}

func (h *minheap) Delete(i int) {
	if _, found := h.pos[i]; !found {
		return
	}
	heap.Remove(h, h.pos[i])
}

func (h minheap) Len() int { return len(h.idx) }

func (h minheap) Less(i, j int) bool { return h.sgn*h.nums[h.idx[i]] < h.sgn*h.nums[h.idx[j]] }

func (h *minheap) Swap(i, j int) {
	h.idx[i], h.idx[j] = h.idx[j], h.idx[i]
	h.pos[h.idx[i]] = i
	h.pos[h.idx[j]] = j
}

func (h *minheap) Push(x interface{}) {
	i := x.(int)
	h.idx = append(h.idx, i)
	h.pos[i] = len(h.idx) - 1
}

func (h *minheap) Pop() interface{} {
	top := len(h.idx) - 1
	i := h.idx[top]
	h.idx = h.idx[:top]
	delete(h.pos, i)
	return i
}
