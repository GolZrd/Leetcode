package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	res := 0
	lowest := prices[0]
	for _, price := range prices {
		if price < lowest {
			lowest = price
		}
		if price-lowest > res {
			res = price - lowest
		}

	}
	return res
}

// Решение с двумя указателями
// func maxProfit(prices []int) int {
// 	res := 0
// 	l, r := 0, 1
// 	for r < len(prices) {
// 		if prices[l] < prices[r] {
// 			profit := prices[r] - prices[l]
// 			if profit > res {
// 				res = profit
// 			}
// 		} else {
// 			l = r
// 		}
// 		r++
// 	}
// 	return res
// }
