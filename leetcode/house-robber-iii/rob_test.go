package rob

import "testing"

func TestRobHouses(t *testing.T) {
	cases := []struct {
		Root   *TreeNode
		Expect int
	}{
		{
			Root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 1},
				},
			},
			Expect: 7,
		},
		{
			Root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 1},
				},
			},
			Expect: 9,
		},
	}
	for _, c := range cases {
		got := rob(c.Root)
		if got != c.Expect {
			t.Errorf("got %d; expect %d", got, c.Expect)
		}
	}
}
