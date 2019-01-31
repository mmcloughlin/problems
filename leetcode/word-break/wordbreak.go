package wordbreak

import "strings"

func wordBreak(s string, wordDict []string) bool {
	g := graph(s, wordDict)
	reachable := map[int]bool{
		0: true,
	}
	stack := []int{0}

	for len(stack) > 0 {
		top := len(stack) - 1
		i := stack[top]
		stack = stack[:top]

		if i == len(s) {
			return true
		}

		for j := range g[i] {
			if reachable[j] {
				continue
			}
			reachable[j] = true
			stack = append(stack, j)
		}
	}

	return false
}

// graph constructs a graph with an edge (i, j) if s[i:j] is in words.
func graph(s string, words []string) []map[int]bool {
	n := len(s)

	// Start with empty graph.
	g := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		g[i] = map[int]bool{}
	}

	// Populate.
	for _, word := range words {
		for i := 0; i < n; i++ {
			if strings.HasPrefix(s[i:], word) {
				g[i][i+len(word)] = true
			}
		}
	}

	return g
}

/*
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) && wordBreak(s[len(word):], wordDict) {
			return true
		}
	}
	return false
}
*/
