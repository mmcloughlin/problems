package solution

func fourSumCount(A []int, B []int, C []int, D []int) int {
	AB := countpairs(A, B)
	CD := countpairs(C, D)
	c := 0
	for ab, m := range AB {
		c += m * CD[-ab]
	}
	return c
}

func countpairs(X, Y []int) map[int]int {
	c := map[int]int{}
	for _, x := range X {
		for _, y := range Y {
			c[x+y]++
		}
	}
	return c
}
