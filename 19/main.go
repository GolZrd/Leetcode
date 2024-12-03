package main

import "fmt"

func main() {
	node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	fmt.Println(removeNthFromEnd(node, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Решение с двумя указателями two pointers
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	// Определяем левый указатель
	left := dummy
	// Определяем правый указатель, он будет указывать на узел head+n
	right := head
	for n > 0 {
		right = right.Next
		n--
	}

	// Теперь мы в цикле смещаем указатели, пока правый не будет nil
	for right != nil {
		left = left.Next
		right = right.Next
	}

	// Когда правый указатель дошел до конца списка, у нас левый стоит на узле n-1,
	// То есть следующий элемент, это тот который нужно удалить, поэтому у левого указываем через него
	left.Next = left.Next.Next

	// Таким образом мы перешагнули нужный узел и его как бы удалили

	// Теперь мы должны вернуть новый лист, dummy у нас была какая надстройка, которая указывала на head
	// Поэтому мы теперь должны вернуть dummy.Next, что будет равно head.
	// Таким образом мы не хотим добавлять нод с левой стороны, поэтому указываем на следующий

	return dummy.Next
}

// //Решение с счетчиком и 2 проходами
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	curr := head

// 	// Подсчитываем количество узлов
// 	count := 0
// 	for curr != nil {
// 		count++
// 		curr = curr.Next
// 	}

// 	// находим узел по порядку, который нужно удалить
// 	removeInd := count - n
// 	if removeInd == 0 {
// 		return head.Next
// 	}

// 	curr = head

// 	for i := 0; i < count-1; i++ {
// 		if (i + 1) == removeInd {
// 			curr.Next = curr.Next.Next
// 			break
// 		}
// 		curr = curr.Next
// 	}

// 	return head
// }
