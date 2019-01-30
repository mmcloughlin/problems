package zigzaglevel

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	levels := [][]int{}
	level := []*TreeNode{root}
	flip := false
	for len(level) > 0 {
		next := []*TreeNode{}
		vals := []int{}
		for _, node := range level {
			if node == nil {
				continue
			}
			vals = append(vals, node.Val)
			next = append(next, node.Left, node.Right)
		}
		for l, r := 0, len(vals)-1; flip && l < r; l, r = l+1, r-1 {
			vals[l], vals[r] = vals[r], vals[l]
		}
		if len(vals) > 0 {
			levels = append(levels, vals)
		}
		level = next
		flip = !flip
	}
	return levels
}
