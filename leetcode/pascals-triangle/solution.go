package solution

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	rows := [][]int{{1}}
	for r := 1; r < numRows; r++ {
		row := make([]int, r+1)
		for i := -1; i < r; i++ {
			row[i+1] = index(rows[r-1], i) + index(rows[r-1], i+1)
		}
		rows = append(rows, row)
	}
	return rows
}

func index(x []int, i int) int {
	if i < 0 || i >= len(x) {
		return 0
	}
	return x[i]
}
