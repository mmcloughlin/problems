package construct

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// First element of preorder is the root value.
	root := &TreeNode{Val: preorder[0]}

	// Look for it in the inorder traveral.
	var split int
	for i, val := range inorder {
		if val == root.Val {
			split = i
			break
		}
	}

	// Everything to the left is the inrder traversal of the left subtree.
	// Likewise for right.
	root.Left = buildTree(preorder[1:1+split], inorder[:split])
	root.Right = buildTree(preorder[1+split:], inorder[split+1:])

	return root
}
