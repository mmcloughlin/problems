package solution

func isPalindrome(s string) bool {
	b := []byte(s)
	n := len(b)
	l, r := 0, n-1
	for {
		// Step until the next alphanumeric.
		for ; l < n && !isalnum(b[l]); l++ {
		}
		for ; r >= 0 && !isalnum(b[r]); r-- {
		}

		if l >= r {
			return true
		}
		if tolower(b[l]) != tolower(b[r]) {
			return false
		}

		l++
		r--
	}
}

func isalpha(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isalnum(ch byte) bool {
	return isalpha(ch) || ('0' <= ch && ch <= '9')
}

func tolower(ch byte) byte {
	if !isalpha(ch) {
		return ch
	}
	return ch | 0x20
}
