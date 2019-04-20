package solution

import (
	"testing"
)

func TestIsSubtree(t *testing.T) {
	cases := []struct {
		S      *TreeNode
		T      *TreeNode
		Expect bool
	}{
		{
			S: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			T: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			Expect: true,
		},
		{
			S: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 0},
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			T: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			Expect: false,
		},
		{
			S:      &TreeNode{Val: 1},
			T:      nil,
			Expect: true,
		},
	}
	for _, c := range cases {
		if got := isSubtree(c.S, c.T); got != c.Expect {
			t.Errorf("isSubtree(%v, %v) = %v; expect %v", c.S, c.T, got, c.Expect)
		}
	}
}
