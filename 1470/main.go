package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{1, 1, 2, 2}, 2))
}

func shuffle(nums []int, n int) []int {
	res := make([]int, 0, len(nums))

	for i := 0; i < n; i++ {
		res = append(res, nums[i], nums[i+n])
	}

	return res
}
