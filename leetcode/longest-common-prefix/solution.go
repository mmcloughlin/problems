package solution

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix string
	for i, b := range []byte(strs[0]) {
		for _, s := range strs[1:] {
			sb := []byte(s)
			if i >= len(sb) || sb[i] != b {
				return prefix
			}
		}
		prefix += string([]byte{b})
	}
	return prefix
}
