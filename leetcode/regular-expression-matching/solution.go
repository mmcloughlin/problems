package solution

func isMatch(s string, p string) bool {
	switch {
	case len(p) == 0:
		return len(s) == 0
	case len(p) > 1 && p[1] == '*':
		for i := 0; i <= len(s); i++ {
			if isMatch(s[i:], p[2:]) {
				return true
			}
			if p[0] != '.' && i < len(s) && s[i] != p[0] {
				break
			}
		}
		return false
	case p[0] == '.':
		return len(s) > 0 && isMatch(s[1:], p[1:])
	default:
		return len(s) > 0 && s[0] == p[0] && isMatch(s[1:], p[1:])
	}
}
