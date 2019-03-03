package solution

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	assert(t, trie.Search("apple"))
	assert(t, !trie.Search("app"))
	assert(t, trie.StartsWith("app"))
	trie.Insert("app")
	assert(t, trie.Search("app"))
}

func assert(t *testing.T, v bool) {
	if !v {
		t.FailNow()
	}
}
