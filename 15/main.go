package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	res := [][]int{}

	slices.Sort(nums)

	for i, j := range nums {
		if i > 0 && j == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1

		for l < r {
			sum := j + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{j, nums[l], nums[r]})
				l++
				// Проверяем если новый элемент равен предыдущему, то идем до тех пор пока не будет отличаться
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}

	}

	return res
}
