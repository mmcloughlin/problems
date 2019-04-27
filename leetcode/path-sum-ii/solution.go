package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return [][]int{{root.Val}}
		}
		return nil
	}

	lpaths := pathSum(root.Left, sum-root.Val)
	rpaths := pathSum(root.Right, sum-root.Val)

	paths := make([][]int, 0, len(lpaths)+len(rpaths))
	for _, path := range lpaths {
		paths = append(paths, append([]int{root.Val}, path...))
	}
	for _, path := range rpaths {
		paths = append(paths, append([]int{root.Val}, path...))
	}
	return paths
}
