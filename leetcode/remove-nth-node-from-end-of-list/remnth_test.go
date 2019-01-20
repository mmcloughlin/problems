package remnth

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		Input  []int
		N      int
		Expect []int
	}{
		{
			Input:  []int{1, 2, 3, 4, 5},
			N:      2,
			Expect: []int{1, 2, 3, 5},
		},
		{
			Input:  []int{1, 2, 3, 4, 5},
			N:      1,
			Expect: []int{1, 2, 3, 4},
		},
		{
			Input:  []int{1, 2, 3, 4, 5},
			N:      0,
			Expect: []int{1, 2, 3, 4, 5},
		},
		{
			Input:  []int{1, 2, 3, 4, 5},
			N:      5,
			Expect: []int{2, 3, 4, 5},
		},
	}
	for _, c := range cases {
		head := ListFromIntSlice(c.Input)
		head = removeNthFromEnd(head, c.N)
		got := IntSliceFromList(head)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Logf("got=%v", got)
			t.Logf("expect=%v", c.Expect)
			t.FailNow()
		}
	}
}

func ListFromIntSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	return &ListNode{
		Val:  nums[0],
		Next: ListFromIntSlice(nums[1:]),
	}
}

func IntSliceFromList(head *ListNode) []int {
	if head == nil {
		return nil
	}
	rest := IntSliceFromList(head.Next)
	return append([]int{head.Val}, rest...)
}
