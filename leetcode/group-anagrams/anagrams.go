package anagrams

import "sort"

func key(s string) string {
	chars := []byte(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func groupAnagrams(strs []string) [][]string {
	bykey := map[string][]string{}

	for _, str := range strs {
		k := key(str)
		bykey[k] = append(bykey[k], str)
	}

	groups := make([][]string, 0, len(bykey))
	for _, group := range bykey {
		groups = append(groups, group)
	}

	return groups
}
