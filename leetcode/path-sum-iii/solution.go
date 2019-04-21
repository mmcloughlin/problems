package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathSumEndingAt(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathSumEndingAt(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	n := 0
	if root.Val == sum {
		n++
	}
	n += pathSumEndingAt(root.Left, sum-root.Val)
	n += pathSumEndingAt(root.Right, sum-root.Val)
	return n
}
