package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate2([]int{1, 3, 4, 2, 2}))
}

// Решение с помощью сортировки. По времени занимает O(n log n). По памяти все зависит от сортировки
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

// Решение с помощью fast and slow pointers.
// Мы можем представить массив как связный список, где каждый элемент указывает на следующий узел.
// Например, если у нас массив [1, 3, 4, 2, 2], то мы можем представить его как связный список:
// 1 -> 3 -> 2 -> 4 -> 2. То есть значение массива указывает на индекс следующего элемента.
// И если у нас по условию есть дупликат, это значит, что у нас есть цикл. А для того, чтобы найти этот цикл
// Исопльзуется алгоритм Floyd’s Tortoise and Hare.
func findDuplicate2(nums []int) int {
	// Определяем 2 указателя
	slow, fast := 0, 0

	// Находим пересечение, то есть находим цикл с помощью Floyd’s Tortoise and Hare
	for {
		// Для медленного указателя, делаем только 1 переход на следующий узел
		slow = nums[slow]
		// Для быстрого указателя, делаем 2 перехода
		fast = nums[nums[fast]]
		//fmt.Println("slow ", slow, "fast ", fast)
		// Если указатели совпадают, то мы обнаружили цикл и можем выйти
		if slow == fast {
			break
		}
	}

	// Теперь нам нужно найти начало этого цикла
	// Для этого мы определяем новый медленный указатель slow2, он будет работать как и обычный медленный указатель
	slow2 := 0
	for {
		// Теперь оба указателя двигаются с одной скоростью
		// Таким образом мы сможем найти точку, в которой возникает цикл
		slow = nums[slow]
		slow2 = nums[slow2]
		// Как только указатели совпадают, то мы нашли начало цикла
		if slow == slow2 {
			return slow
		}

	}
}
