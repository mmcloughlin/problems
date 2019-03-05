package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	total, _ := maxPathSumTotalAndEndingSelf(root)
	return total.x
}

func maxPathSumTotalAndEndingSelf(root *TreeNode) (opt, opt) {
	if root == nil {
		return empty(), empty()
	}

	lt, ls := maxPathSumTotalAndEndingSelf(root.Left)
	rt, rs := maxPathSumTotalAndEndingSelf(root.Right)

	v := value(root.Val)
	self := max(v.add(ls), v.add(rs), v)
	total := max(lt, rt, self, v.add(ls).add(rs))

	return total, self
}

func max(ns ...opt) opt {
	m := empty()
	for _, n := range ns {
		m = m.max(n)
	}
	return m
}

type opt struct {
	v bool // valid
	x int  // value
}

func empty() opt      { return opt{v: false} }
func value(x int) opt { return opt{v: true, x: x} }

func (a opt) add(b opt) opt {
	if !a.v || !b.v {
		return empty()
	}
	return value(a.x + b.x)
}

func (a opt) max(b opt) opt {
	switch {
	case !a.v:
		return b
	case !b.v:
		return a
	case a.x < b.x:
		return b
	default:
		return a
	}
}
