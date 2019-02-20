package rob

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	i, _ := robIncludeExclude(root)
	return i
}

// robIncludeExclude returns the amount that could be robbed both including and
// excluding itself.
func robIncludeExclude(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	il, el := robIncludeExclude(root.Left)
	ir, er := robIncludeExclude(root.Right)
	return max(root.Val+el+er, il+ir), il + ir
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
