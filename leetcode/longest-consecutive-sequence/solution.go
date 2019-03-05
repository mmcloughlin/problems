package solution

func longestConsecutive(nums []int) int {
	// succ[n] is true if n+1 is in the set
	succ := map[int]bool{}
	for _, num := range nums {
		if _, found := succ[num]; found {
			continue
		}
		_, succ[num] = succ[num+1]
		if _, found := succ[num-1]; found {
			succ[num-1] = true
		}
	}

	// identify the heads of sequences
	heads := map[int]bool{}
	for n := range succ {
		heads[n] = true
	}
	for n, hassucc := range succ {
		if hassucc {
			delete(heads, n+1)
		}
	}

	// count their length
	max := 0
	for h := range heads {
		t := h
		for ; succ[t]; t++ {
		}
		l := t - h + 1
		if l > max {
			max = l
		}
	}

	return max
}
