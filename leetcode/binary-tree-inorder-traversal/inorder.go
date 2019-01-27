package inorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	values := inorderTraversal(root.Left)
	values = append(values, root.Val)
	values = append(values, inorderTraversal(root.Right)...)
	return values
}
