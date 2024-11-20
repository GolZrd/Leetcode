package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}

func searchMatrix(matrix [][]int, target int) bool {
	top, bot := 0, len(matrix)-1
	row_ind := -1

	for top <= bot {
		mid := top + (bot-top)/2
		row := matrix[mid]

		if row[0] > target {
			bot = mid - 1
		} else if row[len(row)-1] < target {
			top = mid + 1
		} else {
			row_ind = mid
			break
		}
	}

	if row_ind == -1 {
		return false
	}

	left, right := 0, len(matrix[row_ind])-1
	for left <= right {
		mid := left + (right-left)/2
		val := matrix[row_ind][mid]
		if val > target {
			right = mid - 1
		} else if val < target {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
