package solution

func countSmaller(nums []int) []int {
	counts := make([]int, len(nums))
	t := &bst{}
	for i := len(nums) - 1; i >= 0; i-- {
		counts[i] = t.countlessthan(nums[i])
		t.insert(nums[i])
	}
	return counts
}

type bst struct {
	size     int  // number of entries in the tree
	occupied bool // whether this node has a value
	val      int  // value at this node
	left     *bst // left subtree: all values <= val
	right    *bst // right subtree: all values > val
}

// insert an element into the tree.
func (t *bst) insert(n int) {
	t.size++
	if !t.occupied {
		t.occupied = true
		t.val = n
		t.left = &bst{}
		t.right = &bst{}
		return
	}
	if n <= t.val {
		t.left.insert(n)
	} else {
		t.right.insert(n)
	}
}

// countlessthan returns the number of elements in t less than n.
func (t *bst) countlessthan(max int) int {
	if !t.occupied {
		return 0
	}
	if t.val >= max {
		return t.left.countlessthan(max)
	}
	return t.left.size + 1 + t.right.countlessthan(max)
}
