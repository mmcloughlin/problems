package solution

func intersect(a, b []int) []int {
	x, y := a, b
	intersection := []int{}
	for len(x) > 0 && len(y) > 0 {
		switch {
		case x[0] < y[0]:
			x = x[1:]
		case x[0] > y[0]:
			y = y[1:]
		default:
			n := len(intersection)
			if n == 0 || intersection[n-1] != x[0] {
				intersection = append(intersection, x[0])
			}
			x = x[1:]
			y = y[1:]
		}
	}
	return intersection
}
