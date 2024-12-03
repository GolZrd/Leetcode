package main

import "fmt"

func main() {
	first := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	second := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	fmt.Println(addTwoNumbers(first, second))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy

	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		total := carry
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}

		carry = total / 10
		head.Next = &ListNode{Val: total % 10}
		head = head.Next
	}
	return dummy.Next
}
