package solution

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	cases := []struct {
		Input  *TreeNode
		Expect []int
	}{
		{
			Input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			Expect: []int{1, 3, 4},
		},
		{
			Input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  6,
						Left: &TreeNode{Val: 7},
					},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			Expect: []int{1, 3, 4, 7},
		},
	}
	for _, c := range cases {
		if got := rightSideView(c.Input); !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("rightSideView(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
