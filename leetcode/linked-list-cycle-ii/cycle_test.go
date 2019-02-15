package cycle

import "testing"

func TestDetectCycle(t *testing.T) {
	cases := []struct {
		Vals     []int
		Pos      int
		HasCycle bool
		Expect   int
	}{
		{[]int{3, 2, 0, -4}, 1, true, 2},
		{[]int{1, 2}, 0, true, 1},
		{[]int{1}, -1, false, 0},
	}
	for _, c := range cases {
		head := rho(c.Vals, c.Pos)
		got := detectCycle(head)
		if c.HasCycle && got.Val != c.Expect {
			t.Errorf("got %d expect %d", got.Val, c.Expect)
		}
		if !c.HasCycle && got != nil {
			t.Error("expected no cycle")
		}
	}
}

func rho(vs []int, pos int) *ListNode {
	n := len(vs)
	ns := make([]*ListNode, n)
	for i, v := range vs {
		ns[i] = &ListNode{Val: v}
	}

	for i := 1; i < n; i++ {
		ns[i-1].Next = ns[i]
	}

	if pos >= 0 {
		ns[n-1].Next = ns[pos]
	}

	return ns[0]
}
