package solution

import "sort"

func getSkyline(buildings [][]int) [][]int {
	// construct start/end "edges"
	type edge struct{ delta, x, h int }
	edges := []edge{}

	for _, b := range buildings {
		l, r, h := b[0], b[1], b[2]
		edges = append(edges, edge{+1, l, h})
		edges = append(edges, edge{-1, r, h})
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].x < edges[j].x
	})

	// process in order and produce keypoints
	k := [][]int{}
	heights := map[int]int{}
	cur := 0
	for i := 0; i < len(edges); {
		s := i

		// update heights counts
		for ; i < len(edges) && edges[i].x == edges[s].x; i++ {
			heights[edges[i].h] += edges[i].delta
		}

		// compute max
		m := 0
		for h, count := range heights {
			if count > 0 && h > m {
				m = h
			}
		}
		if m != cur {
			k = append(k, []int{edges[s].x, m})
			cur = m
		}
	}

	return k
}

/*
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSkyline(buildings [][]int) [][]int {
	// heights[i] is the height of the skyline between [i, i+1)
	heights := []int{}
	for _, building := range buildings {
		l, r, h := building[0], building[1], building[2]

		// Extend heights array if necessary.
		for x := len(heights); x < r; x++ {
			heights = append(heights, 0)
		}

		// Update skyline to account for this building.
		for x := l; x < r; x++ {
			heights[x] = max(heights[x], h)
		}
	}

	// Construct the key points.
	keypoints := [][]int{}
	h := 0
	for x, y := range heights {
		if y != h {
			keypoints = append(keypoints, []int{x, y})
			h = y
		}
	}

	return keypoints
}
*/
