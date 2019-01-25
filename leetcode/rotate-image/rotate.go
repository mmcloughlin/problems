package rotate

func rotate(m [][]int) {
	n := len(m)
	for x := 0; x < (n+1)/2; x++ {
		for y := 0; y < n/2; y++ {
			xp, yp := n-1-x, n-1-y
			t := m[y][x]
			m[y][x] = m[xp][y]
			m[xp][y] = m[yp][xp]
			m[yp][xp] = m[x][yp]
			m[x][yp] = t
		}
	}
}
