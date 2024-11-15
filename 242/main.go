package main

import "fmt"

func main() {
	fmt.Println(isAnagram("rats", "sart"))
}

// Решение используя 1 мапу
func isAnagram(s string, t string) bool {
	m1 := make(map[rune]int)

	for _, i := range s {
		m1[i]++
	}
	for _, j := range t {
		m1[j]--
	}

	for _, v := range m1 {
		if v != 0 {
			return false
		}
	}

	return true
}

// Решение с 2 мапами
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	m1 := make(map[rune]int)
// 	for _, j := range s {
// 		m1[j]++
// 	}
// 	m2 := make(map[rune]int)
// 	for _, j := range t {
// 		m2[j]++
// 	}

// 	for k, v := range m1 {
// 		if v != m2[k] {
// 			return false
// 		}
// 	}
// 	return true
// }

// Решение с массимвом из 26 символов
// func isAnagram(s string, t string) bool {
// 	var arr [26]int

// 	for _, i := range s {
// 		arr[i-'a']++
// 	}
// 	for _, j := range t {
// 		arr[j-'a']--
// 	}

// 	for _, k := range arr {
// 		if k != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
