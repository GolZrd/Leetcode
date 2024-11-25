package main

import "fmt"

func main() {
	node := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}

	// Список: 4->2->1
	fmt.Println(reverseList(node))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, curr, next *ListNode
	prev = nil
	curr = head

	for curr != nil {
		next = curr.Next // сначала сохраняем указатель на следующий узел
		curr.Next = prev // производим разворот, чтобы текущий узел указывал на nil, то есть по сути был последним
		prev = curr      // сдвигаем prev, prev у нас будет теперь текущим узлом
		curr = next      // и теперь сдвигаем curr на следующий элемент
	}
	return prev
}
