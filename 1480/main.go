package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}

// Оптимальное решение
func runningSum(nums []int) []int {
	sum := make([]int, len(nums))

	var prefix_sum int
	for i, j := range nums {
		prefix_sum += j
		sum[i] = prefix_sum
	}

	return sum
}

// Решение 1
// func runningSum(nums []int) []int {
// 	sum := make([]int, len(nums))
// 	sum[0] = nums[0]

// 	for i := 1; i < len(nums); i++ {
// 		n := nums[0]
// 		for j := i; j > 0; j-- {
// 			n += nums[j]
// 		}
// 		sum[i] = n
// 	}

// 	return sum
// }
