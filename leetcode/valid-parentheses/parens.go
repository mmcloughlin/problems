package parens

var matching = map[byte]byte{
	']': '[',
	')': '(',
	'}': '{',
}

func isValid(s string) bool {
	stack := []byte{}
	for _, b := range []byte(s) {
		switch b {
		case '(', '{', '[':
			stack = append(stack, b)
			continue
		}

		top := len(stack) - 1
		if top < 0 {
			return false
		}
		ch := stack[top]
		stack = stack[:top]
		expect := matching[b]
		if ch != expect {
			return false
		}
	}
	return len(stack) == 0
}
