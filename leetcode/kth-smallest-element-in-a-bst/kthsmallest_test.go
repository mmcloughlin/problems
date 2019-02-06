package kthsmallest

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	cases := []struct {
		Tree  *TreeNode
		Order []int
	}{
		{
			Tree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 4},
			},
			Order: []int{1, 2, 3, 4},
		},
		{
			Tree: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 1},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			Order: []int{1, 2, 3, 4, 5, 6},
		},
		{
			Tree:  &TreeNode{Val: 1},
			Order: []int{1},
		},
		{
			Tree: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			Order: []int{1, 2},
		},
	}
	for _, c := range cases {
		for i, expect := range c.Order {
			got := kthSmallest(c.Tree, i+1)
			if got != expect {
				t.Errorf("got %v expect %v", got, expect)
			}
		}
	}
}
