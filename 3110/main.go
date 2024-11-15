package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(scoreOfString("zaz"))
}

// Мое решение
func scoreOfString(s string) int {
	var sum int
	for i := 0; i < len(s)-1; i++ {
		sum += int((math.Abs(float64(s[i]) - (float64(s[i+1])))))
	}
	return sum
}

// Интересное еще решение, меньше памяти требуется
// func scoreOfString(s string) int {
// 	var ans int
// 	for i := 1; i < len(s); i++ {
// 		if s[i] > s[i-1] {
// 			ans += int(s[i] - s[i-1])
// 		} else {
// 			ans += int(s[i-1] - s[i])
// 		}
// 	}
// 	return ans
// }
