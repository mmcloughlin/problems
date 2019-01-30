package spiral

func spiralOrder(matrix [][]int) []int {
	s := []int{}
	if len(matrix) == 0 {
		return s
	}

	t, b := 0, len(matrix)-1
	l, r := 0, len(matrix[0])-1

	for t <= b && l <= r {
		// top
		for i := l; i <= r; i++ {
			s = append(s, matrix[t][i])
		}
		t++

		// right
		for i := t; i <= b; i++ {
			s = append(s, matrix[i][r])
		}
		r--

		if t > b || l > r {
			break
		}

		// bottom
		for i := r; i >= l; i-- {
			s = append(s, matrix[b][i])
		}
		b--

		// left
		for i := b; i >= t; i-- {
			s = append(s, matrix[i][l])
		}
		l++
	}

	return s
}
