package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	words := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	expect := []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}

	got := findAllConcatenatedWordsInADict(words)
	sort.Strings(got)
	sort.Strings(expect)
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("got %v; expect %v", got, expect)
	}
}
