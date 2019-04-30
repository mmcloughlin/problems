package solution

import "fmt"

// CharCounts counts characters in s and returns a human-readable version of the result.
func CharCounts(s string) string {
	counts := make([]int, 256)
	for _, ch := range []byte(s) {
		counts[ch]++
	}

	var out string
	for ch, count := range counts {
		if count == 0 {
			continue
		}
		if len(out) > 0 {
			out += ","
		}
		out += fmt.Sprintf("(%c,%d)", byte(ch), count)
	}
	return out
}
