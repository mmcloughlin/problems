package merge

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	// Build list of interval endpoints: delta +1 for start, -1 for end.
	type endpoint struct {
		x     int
		delta int
	}
	endpoints := make([]endpoint, 0, 2*len(intervals))
	for _, i := range intervals {
		endpoints = append(endpoints, endpoint{i.Start, 1}, endpoint{i.End, -1})
	}

	// Sort endpoints.
	sort.Slice(endpoints, func(i, j int) bool {
		e1 := endpoints[i]
		e2 := endpoints[j]
		return e1.x < e2.x || (e1.x == e2.x && e1.delta > e2.delta)
	})

	// Iterate in order, counting how many intervals we are inside.
	var current Interval
	merged := []Interval{}
	n := 0

	for _, e := range endpoints {
		switch {
		case n == 0 && e.delta == 1:
			current.Start = e.x
		case n == 1 && e.delta == -1:
			current.End = e.x
			merged = append(merged, current)
			current = Interval{}
		}
		n += e.delta
	}

	return merged
}
