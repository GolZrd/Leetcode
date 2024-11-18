package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

// func dailyTemperatures(temperatures []int) []int {
// 	stack := [][]int{}
// 	res := make([]int, len(temperatures))

// 	for index, temp := range temperatures {
// 		for len(stack) != 0 && stack[len(stack)-1][1] < temp {
// 			ind := stack[len(stack)-1][0]
// 			res[ind] = index - ind
// 			stack = stack[:len(stack)-1]
// 		}
// 		stack = append(stack, []int{index, temp})
// 	}

// 	return res
// }

// Решение использующее в стэке только индексы
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}

	for ind, temp := range temperatures {
		for len(stack) != 0 && temp > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			res[index] = ind - index
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ind)
	}
	return res
}
