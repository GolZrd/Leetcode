package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func findMin(nums []int) int {
	min := nums[0]
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < min {
			min = nums[mid]
		}
		if nums[l] < min {
			min = nums[l]
		}

		if nums[l] <= nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return min
}
