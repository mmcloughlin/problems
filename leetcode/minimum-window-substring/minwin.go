package minwin

func minWindow(ss string, ts string) string {
	s := []byte(ss)
	t := []byte(ts)
	n := len(s)

	// Set of target characters.
	target := map[byte]int{}
	for _, b := range t {
		target[b]++
	}

	// Our window is all chars in range [tail, head).
	satisfied := 0
	count := map[byte]int{}
	window := ""
	for tail, head := 0, 0; head < n; {
		// Expand the window to include s[head]
		b := s[head]
		head++
		if _, ok := target[b]; !ok {
			continue
		}
		count[b]++
		if count[b] == target[b] {
			satisfied++
		}

		// Do we have all the characters we need?
		if satisfied != len(target) {
			continue
		}

		// Advance tail until the window is minimal.
		for {
			b := s[tail]
			tail++
			if _, ok := target[b]; !ok {
				continue
			}
			count[b]--
			// Advancing tail would mean we did not have all the characters we need.
			if count[b] < target[b] {
				possible := string(s[tail-1 : head])
				if window == "" || len(possible) < len(window) {
					window = possible
				}
				satisfied--
				break
			}
		}
	}

	return window
}
