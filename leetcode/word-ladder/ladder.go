package ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	neighbors := wordGraph(wordList)

	distance := map[string]int{beginWord: 1}
	queue := []string{beginWord}
	for len(queue) > 0 {
		// Pop queue.
		w := queue[0]
		queue = queue[1:]

		// Did we reach the target?
		if w == endWord {
			return distance[w]
		}

		// Visit neighbors.
		for neighbor := range neighbors[w] {
			if _, seen := distance[neighbor]; seen {
				continue
			}
			distance[neighbor] = distance[w] + 1
			queue = append(queue, neighbor)
		}
	}

	return 0
}

func wordGraph(words []string) map[string]map[string]bool {
	n := len(words[0])

	g := map[string]map[string]bool{}
	for _, word := range words {
		g[word] = map[string]bool{}
	}

	for i := 0; i < n; i++ {
		groups := map[string][]string{}
		for _, word := range words {
			key := word[:i] + word[i+1:]
			groups[key] = append(groups[key], word)
		}
		for _, group := range groups {
			for _, w := range group {
				for _, v := range group {
					if w != v {
						g[w][v] = true
					}
				}
			}
		}
	}

	return g
}
