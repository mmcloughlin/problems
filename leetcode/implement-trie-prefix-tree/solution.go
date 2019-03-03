package solution

type Trie struct {
	exists bool
	sub    map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		exists: false,
		sub:    make(map[byte]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.exists = true
		return
	}
	b, rest := word[0], word[1:]
	if _, ok := this.sub[b]; !ok {
		s := Constructor()
		this.sub[b] = &s
	}
	this.sub[b].Insert(rest)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.exists
	}
	b, rest := word[0], word[1:]
	s, ok := this.sub[b]
	return ok && s.Search(rest)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	b, rest := prefix[0], prefix[1:]
	s, ok := this.sub[b]
	return ok && s.StartsWith(rest)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
