package solution

func mySqrt(x int) int {
	r, s := 1, 1
	for s <= x {
		s += 2*r + 1
		r++
	}
	return r - 1
}
