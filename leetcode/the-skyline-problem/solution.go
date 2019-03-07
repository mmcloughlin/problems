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
