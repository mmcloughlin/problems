package solution

import "strings"

/*
func wordBreak(s string, wordDict []string) []string {
	sentencesAt := [][]string{{""}}
	for i := 1; i <= len(s); i++ {
		sentences := []string{}
		for _, word := range wordDict {
			n := len(word)
			if i < n || s[i-n:i] != word {
				continue
			}
			for _, sentence := range sentencesAt[i-n] {
				sentences = append(sentences, concat(sentence, word))
			}
		}
		sentencesAt = append(sentencesAt, sentences)
	}
	return sentencesAt[len(s)]
}

func concat(sentence, word string) string {
	if sentence == "" {
		return word
	}
	return sentence + " " + word
}
*/

func wordBreak(s string, wordDict []string) []string {
	graph := buildGraph(s, wordDict)

	// check reachability
	reachable := make([]bool, len(s)+1)
	reachable[0] = true
	pos := []int{0}
	for len(pos) > 0 {
		top := len(pos) - 1
		p := pos[top]
		pos = pos[:top]

		for _, j := range graph[p] {
			if reachable[j] {
				continue
			}
			reachable[j] = true
			pos = append(pos, j)
		}
	}

	if !reachable[len(s)] {
		return []string{}
	}

	// construct the list of all possible sentences
	type state struct {
		i     int      // index into s
		words []string // accumulated words
	}
	stack := []state{{}}
	sentences := []string{}
	for len(stack) > 0 {
		top := len(stack) - 1
		c := stack[top]
		stack = stack[:top]

		if c.i == len(s) {
			sentences = append(sentences, strings.Join(c.words, " "))
			continue
		}

		for _, j := range graph[c.i] {
			next := state{i: j}
			next.words = append(next.words, c.words...)
			next.words = append(next.words, s[c.i:j])
			stack = append(stack, next)
		}
	}

	return sentences
}

// buildGraph returns a graph such that for all j in g[i], s[i:j] is a word in
// the dictionary.
func buildGraph(s string, wordDict []string) map[int][]int {
	g := map[int][]int{}
	for _, word := range wordDict {
		n := len(word)
		for i := 0; i+n <= len(s); i++ {
			if word == s[i:i+n] {
				g[i] = append(g[i], i+n)
			}
		}
	}
	return g
}
