package solution

func judgeSquareSum(c int) bool {
	squares := map[int]bool{}
	for n := 0; n*n <= c; n++ {
		squares[n*n] = true
		if squares[c-n*n] {
			return true
		}
	}
	return false
}
