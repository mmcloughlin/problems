package solution

func isAnagram(s string, t string) bool {
	count := map[byte]int{}

	for _, b := range []byte(s) {
		count[b]++
	}

	for _, b := range []byte(t) {
		count[b]--
	}

	for _, n := range count {
		if n != 0 {
			return false
		}
	}

	return true
}
