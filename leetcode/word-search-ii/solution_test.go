package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindWords(t *testing.T) {
	cases := []struct {
		Board  [][]byte
		Words  []string
		Expect []string
	}{
		{
			Board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			Words:  []string{"oath", "pea", "eat", "rain"},
			Expect: []string{"eat", "oath"},
		},
		{
			Board: [][]byte{
				{'a', 'a'},
			},
			Words:  []string{"a"},
			Expect: []string{"a"},
		},
		{
			Board: [][]byte{
				{'a'},
			},
			Words:  []string{"a"},
			Expect: []string{"a"},
		},
	}

	for _, c := range cases {
		got := findWords(c.Board, c.Words)
		sort.Strings(got)
		sort.Strings(c.Expect)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Fatalf("got %v; expect %v", got, c.Expect)
		}
	}
}
