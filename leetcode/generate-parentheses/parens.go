package parens

func generateParenthesis(n int) []string {
	solutions := []string{}

	type state struct {
		Parens string
		Open   int
		Close  int
	}
	stack := []state{{}}

	for len(stack) > 0 {
		// Pop stack.
		top := len(stack) - 1
		s := stack[top]
		stack = stack[:top]

		// Are we done?
		if s.Open == n && s.Close == n {
			solutions = append(solutions, s.Parens)
			continue
		}

		// Can we open another paren?
		if s.Open < n {
			n := s
			n.Parens += "("
			n.Open++
			stack = append(stack, n)
		}

		// Can we close a paren?
		if s.Open > s.Close {
			n := s
			n.Parens += ")"
			n.Close++
			stack = append(stack, n)
		}
	}

	return solutions
}
