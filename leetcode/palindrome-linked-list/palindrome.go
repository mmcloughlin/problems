package palindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// Count length.
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}

	// Step to half way, reversing the first half on the way.
	m := head
	var pre *ListNode
	for i := 0; i < n/2; i++ {
		t := m.Next
		m.Next = pre
		pre = m
		m = t
	}

	// Check for palindrome, reversing on the way back too.
	left := pre
	right := m
	if (n % 2) == 1 {
		right = right.Next
	}

	result := true
	for left != nil && right != nil {
		if left.Val != right.Val {
			result = false
		}

		left = left.Next
		right = right.Next
	}

	return result
}
