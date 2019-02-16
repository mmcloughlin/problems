package decode

func decodeString(s string) string {
	dec, _ := parse([]byte(s))
	return dec
}

func parse(b []byte) (string, []byte) {
	result := ""
	for len(b) > 0 {
		if b[0] == ']' {
			break
		}
		if !isdigit(b[0]) {
			result += string(b[:1])
			b = b[1:]
			continue
		}

		var rep int
		var sub string

		rep, b = parsenumber(b)
		b = expect('[', b)
		sub, b = parse(b)
		b = expect(']', b)

		for i := 0; i < rep; i++ {
			result += sub
		}
	}
	return result, b
}

func parsenumber(b []byte) (int, []byte) {
	n := 0
	for len(b) > 0 && isdigit(b[0]) {
		d := int(b[0] - '0')
		n = 10*n + d
		b = b[1:]
	}
	return n, b
}

func isdigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func until(ch byte, b []byte) (string, []byte) {
	i := 0
	for ; i < len(b) && b[i] != ch; i++ {
	}
	return string(b[:i]), b[i:]
}

func expect(ch byte, b []byte) []byte {
	if len(b) == 0 || b[0] != ch {
		panic("unexpected char")
	}
	return b[1:]
}
