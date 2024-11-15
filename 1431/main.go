package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))

	max := 0
	for _, i := range candies {
		if i > max {
			max = i
		}
	}

	for i := 0; i < len(candies); i++ {
		res[i] = candies[i]+extraCandies >= max
	}
	return res
}
