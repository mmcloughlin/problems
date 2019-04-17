package solution

// pos is a board position
type pos struct {
	r, c int
}

func (p pos) Neighbors() []pos {
	return []pos{
		{p.r + 1, p.c},
		{p.r - 1, p.c},
		{p.r, p.c + 1},
		{p.r, p.c - 1},
	}
}

// grid represents a grid of characters.
type grid struct {
	char [][]byte
}

func (g grid) Height() int {
	return len(g.char)
}

func (g grid) Width() int {
	if g.Height() == 0 {
		return 0
	}
	return len(g.char[0])
}

func (g grid) At(p pos) byte {
	return g.char[p.r][p.c]
}

func (g grid) Contains(p pos) bool {
	return 0 <= p.r && p.r < g.Height() && 0 <= p.c && p.c < g.Width()
}

func (g grid) Neighbors(p pos) []pos {
	ns := make([]pos, 0, 4)
	for _, n := range p.Neighbors() {
		if g.Contains(n) {
			ns = append(ns, n)
		}
	}
	return ns
}

func findWords(board [][]byte, words []string) []string {
	g := grid{char: board}
	t := NewTrieDictionary(words)

	set := map[string]bool{}
	var p pos
	for p.r = 0; p.r < g.Height(); p.r++ {
		for p.c = 0; p.c < g.Width(); p.c++ {
			found := findWordsTrie(g, p, map[pos]bool{}, "", t)
			for w := range found {
				set[w] = true
			}
		}
	}

	result := []string{}
	for word := range set {
		result = append(result, word)
	}

	return result
}

func findWordsTrie(g grid, p pos, seen map[pos]bool, prefix string, t *Trie) map[string]bool {
	found := map[string]bool{}

	// Consider this position.
	ch := g.At(p)
	s := t.Sub(ch)
	if s == nil {
		return found
	}
	prefix += string([]byte{ch})

	// Have we found a word at this point.
	if s.Exists() {
		found[prefix] = true
	}

	// Investigate neighbors.
	seen[p] = true
	for _, n := range g.Neighbors(p) {
		if seen[n] {
			continue
		}
		for w := range findWordsTrie(g, n, seen, prefix, s) {
			found[w] = true
		}
	}
	seen[p] = false

	return found
}

type Trie struct {
	exists bool
	sub    map[byte]*Trie
}

func NewTrie() *Trie {
	return &Trie{
		exists: false,
		sub:    make(map[byte]*Trie),
	}
}

func NewTrieDictionary(words []string) *Trie {
	t := NewTrie()
	for _, word := range words {
		t.Insert(word)
	}
	return t
}

func (t *Trie) Exists() bool {
	return t.exists
}

func (t *Trie) Sub(ch byte) *Trie {
	return t.sub[ch]
}

func (t *Trie) Insert(s string) {
	t.InsertBytes([]byte(s))
}

func (t *Trie) InsertBytes(b []byte) {
	if len(b) == 0 {
		t.exists = true
		return
	}
	ch, rest := b[0], b[1:]
	if _, ok := t.sub[ch]; !ok {
		t.sub[ch] = NewTrie()
	}
	t.sub[ch].InsertBytes(rest)
}
