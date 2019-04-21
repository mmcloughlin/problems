package solution

func longestValidParentheses(s string) int {
	longest := 0
	b := []byte(s)
	start := []int{}
	inside := 0
	for i, p := range b {
		// open paren
		if p == '(' {
			if inside >= len(start) {
				start = append(start, i)
			}
			inside++
			continue
		}

		// close must be matched
		if inside < len(start) {
			start = start[:inside]
		}
		if inside == 0 {
			continue
		}

		inside--
		longest = max(i-start[inside]+1, longest)
	}
	return longest
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
