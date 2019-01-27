package sudoku

func isValidSudoku(board [][]byte) bool {
	rowmask := make([]uint16, 9)
	colmask := make([]uint16, 9)
	boxmask := make([]uint16, 9)

	for r, row := range board {
		for c, d := range row {
			if d == '.' {
				continue
			}
			n := uint(d - '0')
			bit := uint16(1) << n

			b := (r / 3) + 3*(c/3)

			if rowmask[r]&bit != 0 {
				return false
			}
			rowmask[r] |= bit

			if colmask[c]&bit != 0 {
				return false
			}
			colmask[c] |= bit

			if boxmask[b]&bit != 0 {
				return false
			}
			boxmask[b] |= bit
		}
	}

	return true
}
