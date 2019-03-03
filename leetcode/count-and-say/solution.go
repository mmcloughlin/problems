package solution

import "strconv"

func countAndSay(n int) string {
	s := "1"
	for i := 2; i <= n; i++ {
		s = rle(s)
	}
	return s
}

// rle run length encodes the string s.
func rle(s string) string {
	e := ""
	b := []byte(s)
	n := len(b)
	for i := 0; i < n; {
		j := i
		for ; j < n && b[j] == b[i]; j++ {
		}
		e += strconv.Itoa(j-i) + string([]byte{b[i]})
		i = j
	}
	return e
}
