package solution

func updateBoard(board [][]byte, click []int) [][]byte {
	g := grid{board}
	b := g.At(click)

	// mine
	if b == 'M' {
		g.Set(click, 'X')
		return board
	}

	// unrevealed empty
	if b == 'E' {
		mines := 0
		for _, a := range g.Adjacent(click) {
			if g.At(a) == 'M' {
				mines++
			}
		}

		if mines > 0 {
			g.Set(click, '0'+byte(mines))
			return board
		}

		g.Set(click, 'B')
		for _, a := range g.Adjacent(click) {
			updateBoard(board, a)
		}
	}

	return board
}

type grid struct {
	board [][]byte
}

func (g grid) Rows() int { return len(g.board) }

func (g grid) Cols() int {
	if g.Rows() == 0 {
		return 0
	}
	return len(g.board[0])
}

func (g grid) At(p []int) byte {
	r, c := p[0], p[1]
	return g.board[r][c]
}

func (g grid) Set(p []int, b byte) {
	r, c := p[0], p[1]
	g.board[r][c] = b
}

func (g grid) Contains(p []int) bool {
	r, c := p[0], p[1]
	return 0 <= r && r < g.Rows() && 0 <= c && c < g.Cols()
}

func (g grid) Adjacent(p []int) [][]int {
	r, c := p[0], p[1]
	pos := [][]int{
		{r - 1, c - 1},
		{r - 1, c},
		{r - 1, c + 1},
		{r, c - 1},
		{r, c},
		{r, c + 1},
		{r + 1, c - 1},
		{r + 1, c},
		{r + 1, c + 1},
	}
	adj := make([][]int, 0, 9)
	for _, p := range pos {
		if g.Contains(p) {
			adj = append(adj, p)
		}
	}
	return adj
}
