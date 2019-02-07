package oddeven

import (
	"reflect"
	"testing"
)

func TestOddEven(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect []int
	}{
		{
			Input:  []int{1, 2, 3, 4, 5},
			Expect: []int{1, 3, 5, 2, 4},
		},
		{
			Input:  []int{2, 1, 3, 5, 6, 4, 7},
			Expect: []int{2, 3, 6, 7, 1, 5, 4},
		},
		{
			Input:  []int{},
			Expect: []int{},
		},
	}
	for _, c := range cases {
		got := ListToSlice(oddEvenList(SliceToList(c.Input)))
		if !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("oddEqual(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}

func ListToSlice(node *ListNode) []int {
	if node == nil {
		return []int{}
	}
	return append([]int{node.Val}, ListToSlice(node.Next)...)
}

func SliceToList(vs []int) *ListNode {
	if len(vs) == 0 {
		return nil
	}
	return &ListNode{
		Val:  vs[0],
		Next: SliceToList(vs[1:]),
	}
}
