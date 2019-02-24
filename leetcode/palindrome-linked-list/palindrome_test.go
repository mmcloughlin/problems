package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect bool
	}{
		{Input: []int{}, Expect: true},
		{Input: []int{1}, Expect: true},
		{Input: []int{1, 2}, Expect: false},
		{Input: []int{1, 2, 3}, Expect: false},
		{Input: []int{1, 2, 1}, Expect: true},
		{Input: []int{1, 2, 2, 1}, Expect: true},
		{Input: []int{1, 2, 3, 1}, Expect: false},
		{Input: []int{1, 2, 2, 3}, Expect: false},
		{Input: []int{1, 2, 3, 2, 1}, Expect: true},
	}
	for _, c := range cases {
		t.Log(c.Input)
		got := isPalindrome(IntsToList(c.Input))
		if got != c.Expect {
			t.Errorf("isPalindrome(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}

func IntsToList(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}
	return &ListNode{
		Val:  is[0],
		Next: IntsToList(is[1:]),
	}
}
