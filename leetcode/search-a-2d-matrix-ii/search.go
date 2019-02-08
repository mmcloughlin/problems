package box

type quad struct {
	l, r, t, b int
}

func (q quad) Width() int  { return q.r - q.l }
func (q quad) Height() int { return q.b - q.t }
func (q quad) Area() int   { return q.Width() * q.Height() }

func (q quad) Sub() []quad {
	mv := (q.t + q.b) / 2
	mh := (q.l + q.r) / 2
	return []quad{
		{q.l, mh, q.t, mv}, // top left
		{mh, q.r, q.t, mv}, // top right
		{q.l, mh, mv, q.b}, // bottom left
		{mh, q.r, mv, q.b}, // bottom right
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	w, h := len(matrix[0]), len(matrix)
	b := quad{0, w, 0, h}
	stack := []quad{b}

	for len(stack) > 0 {
		top := len(stack) - 1
		q := stack[top]
		stack = stack[:top]

		if q.Area() == 0 {
			continue
		}

		if !(matrix[q.t][q.l] <= target && target <= matrix[q.b-1][q.r-1]) {
			continue
		}

		if q.Area() == 1 {
			return true
		}

		stack = append(stack, q.Sub()...)
	}

	return false
}
