package solution

import "fmt"

func largestRectangleArea(heights []int) int {
	a := 0

	// Maintain list of rectangles ending at the left edge.
	rects := []*rect{}

	// Sweep left to right.
	for _, height := range heights {
		fmt.Printf("%v\n", height)

		// All the others can be extended to the right.
		for _, r := range rects {
			fmt.Printf("\tw=%d h=%d a=%d\n", r.w, r.h, r.area())
			r.h = min(r.h, height)
			r.w++
			a = max(a, r.area())
		}

		// This single bar is also a valid rectangle.
		bar := &rect{w: 1, h: height}
		a = max(a, bar.area())
		rects = append(rects, bar)

		// Trim down the list to unique heights only.
		k := 0
		for i := 0; i < len(rects); {
			rects[k] = rects[i]
			j := i + 1
			for j < len(rects) && rects[j].h == rects[i].h {
				rects[k].w = max(rects[k].w, rects[j].w)
				j++
			}
			i = j
			k++
		}
		rects = rects[:k]
	}

	return a
}

type rect struct {
	w, h int
}

func (r rect) area() int {
	return r.w * r.h
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
