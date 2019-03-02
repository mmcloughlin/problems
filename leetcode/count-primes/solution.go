package solution

func countPrimes(n int) int {
	iscomposite := make([]bool, n)
	num := 0
	for p := 2; p < n; p++ {
		if iscomposite[p] {
			continue
		}
		num++
		for m := p * p; m < n; m += p {
			iscomposite[m] = true
		}
	}
	return num
}
