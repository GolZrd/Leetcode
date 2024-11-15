package main

import (
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}

// Используя bucket sort. Time O(n). Space O(n).
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	freq := make([][]int, len(nums)+1)
	for num, count := range counts {
		freq[count] = append(freq[count], num)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

// Используя сортировку
// func topKFrequent(nums []int, k int) []int {
// 	counts := make(map[int]int)
// 	for _, i := range nums {
// 		counts[i]++
// 	}

// 	arr := make([][2]int, 0, len(counts))
// 	for k, v := range counts {
// 		arr = append(arr, [2]int{v, k})
// 	}

// 	// Используем сортировку
// 	sort.Slice(arr, func(i, j int) bool {
// 		return arr[i][0] > arr[j][0]
// 	})

// 	res := make([]int, k)
// 	for i := 0; i < k; i++ {
// 		res[i] = arr[i][1]
// 	}

// 	return res
// }
