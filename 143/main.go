package main

import "fmt"

func main() {
	node := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}

	reorderList(node)

	fmt.Println(node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Сначала нужно определить середину в нашем связном списке
	// Делается это с помощью двух указателей
	slow, fast := head, head.Next

	// Используем цикл для поиска середины списка, пока fast не достигнет конца списка
	for fast != nil && fast.Next != nil {
		// также двигаем slow, это и будет середина списка, а fast это конец списка
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Теперь, когда мы достигли середины списка, то мы можем развернуть вторую половину списка
	second := slow.Next
	// В итоговом списке, у нас средний узел будет последним и будет указывать на nil
	slow.Next = nil

	// Разворачиваем вторую часть списка
	var prev *ListNode
	for second != nil {
		next := second.Next
		second.Next = prev
		prev = second
		second = next
	}

	// Теперь соединяем два списка
	first := head
	second = prev

	// используем цикл пока второй список second не будет nil
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}
