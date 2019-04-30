package solution

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	merged := make([][]int, 0, n)

	i := 0
	for ; i < n && intervals[i][1] < newInterval[0]; i++ {
		merged = append(merged, intervals[i])
	}

	for ; i < n && !disjoint(intervals[i], newInterval); i++ {
		newInterval = intersection(intervals[i], newInterval)
	}

	merged = append(merged, newInterval)

	for ; i < n; i++ {
		merged = append(merged, intervals[i])
	}

	return merged
}

func disjoint(a, b []int) bool {
	return a[1] < b[0] || b[1] < a[0]
}

func intersection(a, b []int) []int {
	return []int{
		min(a[0], b[0]),
		max(a[1], b[1]),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
