package solution

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	// Determine which account every email is in.
	emailaccount := map[string]int{}
	g := newgroups()
	for i, account := range accounts {
		for _, email := range account[1:] {
			if j, ok := emailaccount[email]; ok {
				g.merge(i, j)
			} else {
				emailaccount[email] = i
			}
		}
	}

	// Rebuild the list.
	merged := [][]string{}
	idx := map[int]int{}
	for email, id := range emailaccount {
		id = g.group(id)
		if _, seen := idx[id]; !seen {
			idx[id] = len(merged)
			name := accounts[id][0]
			merged = append(merged, []string{name})
		}
		i := idx[id]
		merged[i] = append(merged[i], email)
	}

	// Sort the emails.
	for _, account := range merged {
		sort.Strings(account[1:])
	}

	return merged
}

type groups struct {
	link map[int]int
}

func newgroups() *groups {
	return &groups{
		link: map[int]int{},
	}
}

func (g *groups) merge(i, j int) {
	i = g.group(i)
	j = g.group(j)
	switch {
	case i < j:
		g.link[j] = i
	case j < i:
		g.link[i] = j
	}
}

func (g *groups) group(i int) int {
	for {
		if j, ok := g.link[i]; ok {
			i = j
		} else {
			return i
		}
	}
}
