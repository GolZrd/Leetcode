package main

import (
	"fmt"
)

func main() {
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

// Решение 1
func numJewelsInStones(jewels string, stones string) int {
	var count int

	m := make(map[rune]struct{})

	for _, j := range jewels {
		m[j] = struct{}{}
	}

	for _, k := range stones {
		if _, ok := m[k]; ok {
			count++
		}
	}

	return count
}

// Решение с использованием стороной библиотеки
// func numJewelsInStones(jewels string, stones string) int {
// 	var count int

// 	for _, j := range stones {
// 		if strings.Contains(jewels, string(j)) {
// 			count++
// 		}
// 	}
// 	return count
// }

// Еще одно решение
// func numJewelsInStones(jewels string, stones string) int {
// 	var sum int
// 	for i := 0; i < len(jewels); i++ {
// 		for j := 0; j < len(stones); j++ {
// 			if jewels[i] == stones[j] {
// 				sum++
// 			}
// 		}
// 	}
// 	return sum
// }
