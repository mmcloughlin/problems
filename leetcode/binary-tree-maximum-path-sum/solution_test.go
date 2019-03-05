package solution

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	cases := []struct {
		Root   *TreeNode
		Expect int
	}{
		{
			Root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			Expect: 6,
		},
		{
			Root: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			Expect: 42,
		},
		{
			Root:   &TreeNode{Val: -3},
			Expect: -3,
		},
	}
	for _, c := range cases {
		if got := maxPathSum(c.Root); got != c.Expect {
			t.Errorf("maxPathSum() = %v; expect %v", got, c.Expect)
		}
	}
}
