package solution

import "math"

// l[i] = MAX l[n] OVER (neighbors n with m[n] > m[i])

const unknown = math.MinInt32

// longestIncreasingPath computes the length of the longest increasing path in
// matrix.
func longestIncreasingPath(matrix [][]int) int {
	longest := 0
	m := Matrix(matrix)
	l := NewMatrixConst(m.Width(), m.Height(), unknown)
	for r := 0; r < m.Height(); r++ {
		for c := 0; c < m.Width(); c++ {
			d := longestIncreasingPathAt(m, l, Pos{r, c})
			longest = max(d, longest)
		}
	}
	return longest
}

// longestIncreasingPathAt computes the length of the longest increasing path
// starting at p.
func longestIncreasingPathAt(m, l Matrix, p Pos) int {
	if l.At(p) != unknown {
		return l.At(p)
	}

	longest := 1
	for _, n := range m.Neighbors(p) {
		if m.At(n) > m.At(p) {
			d := longestIncreasingPathAt(m, l, n)
			longest = max(longest, d+1)
		}
	}

	l.Set(p, longest)
	return longest
}

type Pos struct {
	r, c int
}

func (p Pos) Neighbors() []Pos {
	return []Pos{
		{p.r + 1, p.c},
		{p.r - 1, p.c},
		{p.r, p.c + 1},
		{p.r, p.c - 1},
	}
}

type Matrix [][]int

func NewMatrix(w, h int) Matrix {
	m := make([][]int, h)
	for r := 0; r < h; r++ {
		m[r] = make([]int, w)
	}
	return m
}

func NewMatrixConst(w, h, v int) Matrix {
	m := NewMatrix(w, h)
	for r := 0; r < m.Height(); r++ {
		for c := 0; c < m.Width(); c++ {
			m[r][c] = v
		}
	}
	return m
}

func (m Matrix) Height() int {
	return len(m)
}

func (m Matrix) Width() int {
	if m.Height() == 0 {
		return 0
	}
	return len(m[0])
}

func (m Matrix) At(p Pos) int {
	return m[p.r][p.c]
}

func (m Matrix) Set(p Pos, v int) {
	m[p.r][p.c] = v
}

func (m Matrix) Contains(p Pos) bool {
	return 0 <= p.r && p.r < m.Height() && 0 <= p.c && p.c < m.Width()
}

func (m Matrix) Neighbors(p Pos) []Pos {
	ns := make([]Pos, 0, 4)
	for _, n := range p.Neighbors() {
		if m.Contains(n) {
			ns = append(ns, n)
		}
	}
	return ns
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
