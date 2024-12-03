package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// Задача решается с помощью Floyd’s Tortoise and Hare algorithm

	// Определяем указатели fast и slow
	fast, slow := head, head

	// Идем по циклу пока fast не равен nil или пока fast.next не равен nil.
	// Используем fast так как он будет идти быстрее
	for fast != nil && fast.Next != nil {
		// внутри цикла сдвигаем slow на следующий узел
		slow = slow.Next
		// А fast свдигаем на два шага (узла)
		fast = fast.Next.Next

		// И теперь нас осталось сделать проверку на равенство slow и fast
		if slow == fast {
			return true
		}
	}

	return false
}
