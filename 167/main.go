package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

//Решение с two pointers
func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			res[0], res[1] = l+1, r+1
			break
		}
	}
	return res
}
