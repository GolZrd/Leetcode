package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {

	max := 0
	left, right := 0, len(height)-1

	for left <= right {
		area := min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

// func maxArea(height []int) int {

// 	max := 0
// 	left, right := 0, len(height)-1

// 	for left < right {
// 		if height[left] < height[right] {
// 			area := height[left] * (right - left)
// 			if area > max {
// 				max = area
// 			}
// 			left++
// 		} else if height[left] > height[right] {
// 			area := height[right] * (right - left)
// 			if area > max {
// 				max = area
// 			}
// 			right--
// 		} else {
// 			area := height[left] * (right - left)
// 			if area > max {
// 				max = area
// 			}
// 			left++
// 		}
// 	}
// 	return max
// }
