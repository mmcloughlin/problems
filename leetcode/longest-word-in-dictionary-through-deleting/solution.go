package solution

import "strings"

func findLongestWord(s string, d []string) string {
	result := ""
	for _, w := range d {
		if !deleteable([]byte(s), []byte(w)) {
			continue
		}
		if len(w) > len(result) || (len(w) == len(result) && strings.Compare(w, result) < 0) {
			result = w
		}
	}
	return result
}

func deleteable(s, w []byte) bool {
	if len(w) == 0 {
		return true
	}
	for i := range s {
		if s[i] == w[0] {
			return deleteable(s[i+1:], w[1:])
		}
	}
	return false
}
