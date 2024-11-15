package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}

// Решение с 2 циклами и срезом
// func twoSum(nums []int, target int) []int {
// 	res := make([]int, 2)
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				res[0], res[1] = i, j
// 				break
// 			}
// 		}
// 	}
// 	return res
// }

// // Более оптимальное решение с помощью мапы
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, j := range nums {
		if v, ok := m[target-j]; ok {
			return []int{v, i}
		} else {
			m[j] = i
		}
	}

	return []int{}
}
