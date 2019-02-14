package flatten

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 6},
		},
	}
	t.Logf(root.String())
	flatten(root)
	t.Logf(root.String())
}

func (n *TreeNode) String() string {
	if n == nil {
		return "<nil>"
	}
	return fmt.Sprintf("(%s %d %s)", n.Left, n.Val, n.Right)
}
