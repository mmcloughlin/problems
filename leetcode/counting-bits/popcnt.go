package popcnt

func countBits(num int) []int {
	p := []int{0}
	for len(p) <= num {
		n := len(p)
		for i := 0; i < n; i++ {
			p = append(p, p[i]+1)
		}
	}
	return p[:num+1]
}
