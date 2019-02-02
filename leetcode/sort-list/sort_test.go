package sort

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortList(t *testing.T) {
	cases := [][]int{
		{4, 2, 1, 3},
		{-1, 5, 3, 4, 0},
		{1},
		{1, 2, 3},
		{},
	}
	for _, c := range cases {
		got := ListToSlice(sortList(SliceToList(c)))
		sort.Ints(c)
		if !reflect.DeepEqual(got, c) {
			t.Errorf("\n   got %v\nexpect %v", got, c)
		}
	}
}

func SliceToList(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}
	return &ListNode{
		Val:  is[0],
		Next: SliceToList(is[1:]),
	}
}

func ListToSlice(node *ListNode) []int {
	if node == nil {
		return []int{}
	}
	return append([]int{node.Val}, ListToSlice(node.Next)...)
}
