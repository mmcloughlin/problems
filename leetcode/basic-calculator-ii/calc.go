package calc

// calculate evaluates the expression s.
func calculate(s string) int {
	mul := binparser(number, map[byte]binop{
		'*': func(a, b int) int { return a * b },
		'/': func(a, b int) int { return a / b },
	})
	p := binparser(mul, map[byte]binop{
		'+': func(a, b int) int { return a + b },
		'-': func(a, b int) int { return a - b },
	})
	n, _ := p([]byte(s))
	return n
}

// binop is a binary operator.
type binop func(int, int) int

// parser parses a number from an expression, and returns the rest.
type parser func([]byte) (int, []byte)

// binparser builds a parser for the given binary operators.
func binparser(sub parser, ops map[byte]binop) parser {
	return func(s []byte) (int, []byte) {
		a, s := sub(s)
		for len(s) > 0 && ops[s[0]] != nil {
			op := ops[s[0]]
			b, t := sub(s[1:])
			a = op(a, b)
			s = t
		}
		return a, s
	}
}

// number parses a number from s.
func number(s []byte) (int, []byte) {
	// discard leading whitespace
	s = chomp(s)

	// parse the number
	n := 0
	for len(s) > 0 && isdigit(s[0]) {
		d := s[0] - '0'
		n = 10*n + int(d)
		s = s[1:]
	}

	// discard trailing whitespace
	s = chomp(s)

	return n, s
}

// chomp consumes whitespace.
func chomp(s []byte) []byte {
	for len(s) > 0 && isspace(s[0]) {
		s = s[1:]
	}
	return s
}

// isspace reports whether b is a whitespace character.
func isspace(b byte) bool {
	return b == ' ' || b == '\t'
}

// isdigit reports whether b is a digit.
func isdigit(b byte) bool {
	return '0' <= b && b <= '9'
}
