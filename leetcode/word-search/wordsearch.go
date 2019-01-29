package wordsearch

func exist(board [][]byte, word string) bool {
	letters := []byte(word)
	for y, row := range board {
		for x := range row {
			p := Pos{x, y}
			if search(p, board, letters, make(map[Pos]bool)) {
				return true
			}
		}
	}
	return false
}

func search(origin Pos, board [][]byte, letters []byte, visited map[Pos]bool) bool {
	if len(letters) == 0 {
		return true
	}

	first, rest := letters[0], letters[1:]
	b := Board{Grid: board}

	if first != b.At(origin) {
		return false
	}

	if len(rest) == 0 {
		return true
	}

	visited[origin] = true
	defer delete(visited, origin)

	for _, n := range b.Neighbors(origin) {
		if visited[n] {
			continue
		}
		if search(n, board, rest, visited) {
			return true
		}
	}

	return false
}

type Pos struct {
	X, Y int
}

func (p Pos) Neighbors() []Pos {
	return []Pos{
		Pos{p.X, p.Y - 1},
		Pos{p.X, p.Y + 1},
		Pos{p.X - 1, p.Y},
		Pos{p.X + 1, p.Y},
	}
}

type Board struct {
	Grid [][]byte
}

func (b Board) At(p Pos) byte {
	return b.Grid[p.Y][p.X]
}

func (b Board) Contains(p Pos) bool {
	return (0 <= p.X && p.X < b.Width()) && (0 <= p.Y && p.Y < b.Height())
}

func (b Board) Width() int {
	if len(b.Grid) == 0 {
		return 0
	}
	return len(b.Grid[0])
}

func (b Board) Height() int {
	return len(b.Grid)
}

func (b Board) Neighbors(p Pos) []Pos {
	ns := make([]Pos, 0, 4)
	for _, n := range p.Neighbors() {
		if b.Contains(n) {
			ns = append(ns, n)
		}
	}
	return ns
}
