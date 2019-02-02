package sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// Quicksort. Pick first value as the pivot.
	pivot := head
	node := head.Next

	// Split list by pivot.
	var left, right *ListNode
	for node != nil {
		next := node.Next
		if node.Val <= pivot.Val {
			node.Next = left
			left = node
		} else {
			node.Next = right
			right = node
		}
		node = next
	}

	// Sort each side.
	left = sortList(left)
	right = sortList(right)

	// Join.
	pivot.Next = right
	if left == nil {
		return pivot
	}
	head = left

	// This part would be much faster if we had a function that returned the
	// tail as well.
	tail := left
	for ; tail.Next != nil; tail = tail.Next {
	}
	tail.Next = pivot
	return head
}
