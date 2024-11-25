package main

import "fmt"

func main() {
	// Список: 1->2->4
	first := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	// Список: 1->3->4
	second := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	fmt.Println(mergeTwoLists(first, second)) // Список: 1->1->2->3->4->4
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// создадим новый связный список
	newListNode := &ListNode{}
	// Чтобы смещаться по новому связному списку и не потерять его начало
	// мы будем использовать переменную, которой присвоим значение newListNode
	head := newListNode

	// Используем цикл for пока один из списков не станет пустым
	for list1 != nil && list2 != nil {
		// Если значение певрого списка меньше, чем второе
		if list1.Val < list2.Val {
			//то вставляем значение list1 в новый список
			head.Next = list1
			// смещаем list1 на следующий узел
			list1 = list1.Next
		} else {
			// иначе вставляем значение list2 в новый список
			head.Next = list2
			// смещаем list2 на следующий узел
			list2 = list2.Next
		}

		// смещаем nodes на следующий узел
		head = head.Next
	}
	//Проверяем какой лист пустой, и вставляем оставшийся в новый список
	if list1 == nil {
		head.Next = list2
	} else {
		head.Next = list1
	}

	return newListNode.Next
}
