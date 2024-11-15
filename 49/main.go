package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// Решение с рунами
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		x := []rune{}
		for _, j := range str {
			x = append(x, j)
		}
		sort.Slice(x, func(i, j int) bool {
			return x[i] < x[j]
		})
		m[string(x)] = append(m[string(x)], str)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// Решение с байтами
// func groupAnagrams(strs []string) [][]string {
// 	m := make(map[string][]string)

// 	for _, str := range strs {
// 		bytes := []byte(str)
// 		slices.Sort(bytes)
// 		sortedSt := string(bytes)

// 		if group, ok := m[sortedSt]; ok {
// 			group = append(group, str)
// 			m[sortedSt] = group
// 		} else {
// 			m[sortedSt] = []string{str}
// 		}
// 	}
// 	res := make([][]string, 0, len(m))
// 	for _, v := range m {
// 		res = append(res, v)
// 	}
// 	return res
// }

// Решение с массивом из 26 символов
// func groupAnagrams(strs []string) [][]string {
// 	m := make(map[[26]int][]string)

// 	for _, str := range strs {
// 		arr := [26]int{}
// 		for _, j := range str {
// 			arr[j-'a']++
// 		}
// 		if _, ok := m[arr]; !ok {
// 			m[arr] = []string{str}
// 		} else {
// 			m[arr] = append(m[arr], str)
// 		}
// 	}
// 	res := make([][]string, 0, len(m))
// 	for _, v := range m {
// 		res = append(res, v)
// 	}
// 	return res
// }
