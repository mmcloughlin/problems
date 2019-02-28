package solution

import "math"

const impossible = -1

func coinChange(coins []int, amount int) int {
	m := []int{0}
	for a := 1; a <= amount; a++ {
		num := math.MaxInt64
		possible := false
		for _, c := range coins {
			p := a - c
			if p >= 0 && m[p] != impossible {
				num = min(num, m[p]+1)
				possible = true
			}
		}
		if !possible {
			num = impossible
		}
		m = append(m, num)
	}
	return m[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
