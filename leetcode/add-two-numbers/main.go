package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	next := func(l *ListNode) *ListNode {
		if l != nil {
			return l.Next
		}
		return nil
	}
	val := func(l *ListNode) int {
		if l != nil {
			return l.Val
		}
		return 0
	}

	base := &ListNode{}
	cur := base
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		s := val(l1) + val(l2) + carry

		cur.Next = &ListNode{Val: s % 10}
		carry = s / 10

		cur = cur.Next
		l1 = next(l1)
		l2 = next(l2)
	}
	return base.Next
}

func (l *ListNode) Print() {
	if l != nil {
		fmt.Printf("%d ", l.Val)
		l.Next.Print()
	}
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	s := addTwoNumbers(l1, l2)
	s.Print()
}
