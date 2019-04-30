package solution

// Interval is the closed interval [Start, End].
type Interval struct {
	Start int
	End   int
}

// Union merges m into the disjoint intervals is.
func Union(is []Interval, m Interval) []Interval {
	n := len(is)

	// Empty input.
	if n == 0 {
		return []Interval{m}
	}

	// m occurs before everything.
	if m.End < is[0].Start {
		return append([]Interval{m}, is...)
	}

	// m occurs after everything.
	if is[n-1].End < m.Start {
		return append(is, m)
	}

	// Now we're in the case where m occurs somewhere in the middle.
	merged := make([]Interval, 0, n)

	i := 0
	for ; is[i].End < m.Start; i++ {
		merged = append(merged, is[i])
	}

	j := i
	for ; j+1 < n && is[j+1].Start <= m.End; j++ {

	}

	// m.Start <= is[i].End
	// is[j].Start <= m.End

	merged = append(merged, Interval{
		Start: min(m.Start, is[i].Start),
		End:   max(m.End, is[j].End),
	})

	for i := j + 1; i < n; i++ {
		merged = append(merged, is[i])
	}

	return merged
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
