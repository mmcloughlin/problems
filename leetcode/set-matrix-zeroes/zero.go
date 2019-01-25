package zero

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	// We will use the first row for storage.
	// First determine if it already has zeros in it.
	firstrowzero := false
	for _, val := range matrix[0] {
		if val == 0 {
			firstrowzero = true
		}
	}

	// Determine which columns must be zero.
	for _, row := range matrix {
		for x, val := range row {
			if val == 0 {
				matrix[0][x] = 0
			}
		}
	}

	for _, row := range matrix[1:] {
		clear := false
		for i := range row {
			if row[i] == 0 {
				clear = true
				break
			}
			if matrix[0][i] == 0 {
				row[i] = 0
			}
		}
		if clear {
			for i := range row {
				row[i] = 0
			}
		}
	}

	if firstrowzero {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
}
