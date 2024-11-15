package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

// Оптимальное решение
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = 1
	}

	pref := 1
	for i := 0; i < len(nums); i++ {
		res[i] = pref
		pref *= nums[i]
	}

	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= post
		post *= nums[i]
	}

	return res
}

// Неоптимальное решение
// func productExceptSelf(nums []int) []int {
// 	res := make([]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		prod := 1
// 		for j := 0; j < len(nums); j++ {
// 			if i == j {
// 				continue
// 			} else {
// 				prod *= nums[j]
// 			}
// 		}
// 		res[i] = prod
// 	}

// 	return res
// }
