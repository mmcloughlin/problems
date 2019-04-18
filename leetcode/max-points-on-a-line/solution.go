package solution

import "sort"

func maxPoints(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}

	// Sort by x-coordinate.
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// Maintain list of lines as we "sweep" left-to-right.
	var lines []*line

	for _, coords := range points {
		p := &point{x: coords[0], y: coords[1]}
		n := []*line{
			{a: p, num: 1},
		}
		for _, l := range lines {
			if l.contains(p) {
				l.num++
			}

			if l.b == nil && *l.a != *p {
				n = append(n, &line{a: l.a, b: p, num: l.num + 1})
			}
		}
		lines = append(lines, n...)
	}

	// Find the max.
	m := 0
	for _, l := range lines {
		m = max(m, l.num)
	}

	return m
}

type point struct {
	x, y int
}

type line struct {
	a   *point
	b   *point
	num int
}

func (l line) contains(p *point) bool {
	if l.b == nil {
		return *l.a == *p
	}

	abx := l.a.x - l.b.x
	aby := l.a.y - l.b.y
	bpx := l.b.x - p.x
	bpy := l.b.y - p.y
	return aby*bpx == bpy*abx
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
