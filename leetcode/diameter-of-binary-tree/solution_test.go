package solution

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	cases := []struct {
		Root   *TreeNode
		Expect int
	}{
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
			Expect: 3,
		},
		{
			Root:   nil,
			Expect: 0,
		},
		{
			Root:   &TreeNode{Val: 1},
			Expect: 0,
		},
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 5},
				},
			},
			Expect: 2,
		},
	}
	for _, c := range cases {
		if got := diameterOfBinaryTree(c.Root); got != c.Expect {
			t.Errorf("got %d; expect %d", got, c.Expect)
		}
	}
}
