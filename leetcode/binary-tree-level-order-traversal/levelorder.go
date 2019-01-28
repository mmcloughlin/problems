package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levels := [][]int{
		{root.Val},
	}

	left := levelOrder(root.Left)
	levels = append(levels, left...)

	right := levelOrder(root.Right)
	for i, values := range right {
		level := i + 1
		if level == len(levels) {
			levels = append(levels, nil)
		}
		levels[level] = append(levels[level], values...)
	}

	return levels
}
