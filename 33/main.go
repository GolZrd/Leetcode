package main

import "fmt"

func main() {
	fmt.Println(search([]int{3, 1}, 1))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			// Это означает что левая сторона отсортирована
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// Если упали сюда, значит правая сторона отсортирована
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
