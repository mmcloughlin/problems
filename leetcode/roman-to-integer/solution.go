package solution

var value = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	n := 0
	p := 0
	for _, b := range []byte(s) {
		v := value[b]
		n += v
		if v > p {
			n -= 2 * p
		}
		p = v
	}
	return n
}
