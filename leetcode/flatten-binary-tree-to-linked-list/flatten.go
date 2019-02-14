package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// flatten flattens root.
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	tail := root.Right

	n := root.Left
	if n == nil {
		return
	}

	for n.Right != nil {
		n = n.Right
	}
	n.Right = tail

	root.Right = root.Left
	root.Left = nil
}
