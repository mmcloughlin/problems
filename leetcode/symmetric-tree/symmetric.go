package symmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || ismirror(root.Left, root.Right)
}

func ismirror(left, right *TreeNode) bool {
	// nil check
	switch {
	case left == nil && right == nil:
		return true
	case left == nil:
		return false
	case right == nil:
		return false
	}

	// Values should match.
	if left.Val != right.Val {
		return false
	}

	return ismirror(left.Left, right.Right) && ismirror(left.Right, right.Left)
}
