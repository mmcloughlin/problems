package solution

import "container/heap"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	w := &window{
		nums:     nums,
		indicies: []int{},
		pos:      map[int]int{},
	}

	// Fill with initial window.
	i := 0
	for ; i < k; i++ {
		heap.Push(w, i)
	}

	// Now push and pop elements.
	max := []int{}
	for ; i < len(nums); i++ {
		max = append(max, w.Max())
		w.Add(i)
		w.Remove(i - k)
	}
	max = append(max, w.Max())

	return max
}

type window struct {
	nums     []int
	indicies []int
	pos      map[int]int
}

func (w *window) Max() int       { return w.nums[w.indicies[0]] }
func (w *window) Remove(idx int) { heap.Remove(w, w.pos[idx]) }
func (w *window) Add(idx int)    { heap.Push(w, idx) }

func (w *window) Len() int           { return len(w.indicies) }
func (w *window) Less(i, j int) bool { return w.nums[w.indicies[i]] > w.nums[w.indicies[j]] }

func (w *window) Swap(i, j int) {
	w.indicies[i], w.indicies[j] = w.indicies[j], w.indicies[i]
	w.pos[w.indicies[i]] = i
	w.pos[w.indicies[j]] = j
}

func (w *window) Push(x interface{}) {
	idx := x.(int)
	n := w.Len()
	w.indicies = append(w.indicies, idx)
	w.pos[idx] = n
}

func (w *window) Pop() interface{} {
	n := w.Len() - 1
	idx := w.indicies[n]
	w.indicies = w.indicies[:n]
	delete(w.pos, idx)
	return idx
}
