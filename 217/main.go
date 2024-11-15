package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
}

// Решение с помощью мапы
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))

	for _, i := range nums {
		if _, ok := m[i]; ok {
			return true
		} else {
			m[i] = struct{}{}
		}
	}

	return false
}
