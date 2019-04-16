package solution

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestMergeKListsExample(t *testing.T) {
	lists := []*ListNode{
		SliceToList([]int{1, 4, 5}),
		SliceToList([]int{1, 3, 4}),
		SliceToList([]int{2, 6}),
	}
	expect := []int{1, 1, 2, 3, 4, 4, 5, 6}
	got := ListToSlice(mergeKLists(lists))
	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("mergeKLists(...) = %v; expect %v", got, expect)
	}
}

func TestMergeKListsRandom(t *testing.T) {
	for trial := 0; trial < 512; trial++ {
		n := 1 + rand.Intn(32)
		lists := make([]*ListNode, n)
		expect := []int{}
		for i := 0; i < n; i++ {
			vals := RandSortedSlice(128, 16)
			lists[i] = SliceToList(vals)
			expect = append(expect, vals...)
		}

		sort.Ints(expect)
		got := ListToSlice(mergeKLists(lists))
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("mergeKLists(...) = %v; expect %v", got, expect)
		}
	}
}

func RandSortedSlice(maxLen, maxDelta int) []int {
	n := rand.Intn(maxLen + 1)
	vals := []int{}
	val := 1
	for i := 0; i < n; i++ {
		val += rand.Intn(maxDelta + 1)
		vals = append(vals, val)
	}
	return vals
}

func SliceToList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	return &ListNode{
		Val:  vals[0],
		Next: SliceToList(vals[1:]),
	}
}

func ListToSlice(n *ListNode) []int {
	if n == nil {
		return nil
	}
	rest := ListToSlice(n.Next)
	return append([]int{n.Val}, rest...)
}
