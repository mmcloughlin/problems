package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	greater, _ := convertGreater(root, 0)
	return greater
}

func convertGreater(root *TreeNode, s int) (*TreeNode, int) {
	if root == nil {
		return nil, s
	}
	r, sr := convertGreater(root.Right, s)
	l, sl := convertGreater(root.Left, sr+root.Val)
	return &TreeNode{
		Val:   root.Val + sr,
		Left:  l,
		Right: r,
	}, sl
}
