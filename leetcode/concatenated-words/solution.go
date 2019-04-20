package solution

func findAllConcatenatedWordsInADict(words []string) []string {
	// Construct words set.
	dict := map[string]bool{}
	for _, word := range words {
		dict[word] = true
	}

	// Check
	concat := map[string]bool{}
	for _, word := range words {
		concat[word] = isConcatenated(word, dict, concat)
	}

	// Return list.
	result := []string{}
	for _, word := range words {
		if concat[word] {
			result = append(result, word)
		}
	}

	return result
}

func isConcatenated(word string, dict, concat map[string]bool) bool {
	// Exit if the result is known.
	if _, found := concat[word]; found {
		return concat[word]
	}

	// Look for prefixes in the dictionary.
	result := false
	for i := 1; i < len(word); i++ {
		prefix, rest := word[:i], word[i:]
		if !dict[prefix] {
			continue
		}
		if dict[rest] || isConcatenated(rest, dict, concat) {
			result = true
			break
		}
	}

	concat[word] = result
	return result
}
