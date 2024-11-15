package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

// func longestConsecutive(nums []int) int {
// 	m := make(map[int]struct{})
// 	for _, i := range nums {
// 		m[i] = struct{}{}
// 	}
// 	res := 0
// 	for _, num := range nums {
// 		if _, ok := m[num-1]; !ok {
// 			leng := 1
// 			for {
// 				if _, exist := m[num+leng]; exist {
// 					leng++
// 				} else {
// 					break
// 				}
// 			}
// 			if leng > res {
// 				res = leng
// 			}
// 		}
// 	}

// 	return res
// }

// Если итерироваться по карте, то задержка будет существеннее меньше чем по слайсу
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, i := range nums {
		m[i] = struct{}{}
	}
	res := 0
	for num := range m {
		if _, ok := m[num-1]; !ok {
			leng := 1
			for {
				if _, exist := m[num+leng]; exist {
					leng++
				} else {
					break
				}
			}
			if leng > res {
				res = leng
			}
		}
	}

	return res
}
